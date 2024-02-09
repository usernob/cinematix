<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import Dropdown from 'flowbite-svelte/Dropdown.svelte';
	import DropdownHeader from 'flowbite-svelte/DropdownHeader.svelte';
	import DropdownItem from 'flowbite-svelte/DropdownItem.svelte';
	import DropdownDivider from 'flowbite-svelte/DropdownDivider.svelte';
	import ArrowLeftSolid from 'flowbite-svelte-icons/ArrowLeftSolid.svelte';
	import HomeOutline from 'flowbite-svelte-icons/HomeOutline.svelte';
	import { goto } from '$app/navigation';
	import type { User } from '@/lib/types/modelTypes';

	export let withArrowBack: boolean = true;
	export let withHome: boolean = false;
	export let title: string = 'Films';
	export let user: User | null = null;
</script>

<Navbar class="fixed top-0 z-50 px-0">
	<div class="flex items-center justify-start gap-4">
		{#if withArrowBack}
			<ArrowLeftSolid class="h-4 w-4" on:click={() => history.back()} />
		{/if}
		{#if withHome}
			<HomeOutline class="h-4 w-4" on:click={() => goto('/')} />
		{/if}
	</div>
	<div class="font-bold md:text-2xl">
		<h3>{title}</h3>
	</div>
	<div class="flex md:order-2">
		<DarkMode class="me-1 md:me-2" />
		{#if user}
			<div class="flex items-center md:order-2">
				<p id="user-menu" class="cursor-pointer text-sm font-medium">Hi, {user.nama}</p>
			</div>
			<Dropdown placement="bottom" triggeredBy="#user-menu">
				<DropdownHeader>
					<span class="block truncate text-sm font-medium">{user.email}</span>
				</DropdownHeader>
				<DropdownItem>Dashboard</DropdownItem>
				<DropdownItem>Settings</DropdownItem>
				<DropdownDivider />
				<DropdownItem>Log out</DropdownItem>
			</Dropdown>
		{:else}
			<a
				href="/login"
				class="rounded-lg border-2 border-primary-700 px-4 py-2 text-sm font-medium text-primary-700"
				>Sign In</a
			>
		{/if}
	</div>
</Navbar>
