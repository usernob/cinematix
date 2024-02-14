<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import Dropdown from 'flowbite-svelte/Dropdown.svelte';
	import Avatar from 'flowbite-svelte/Avatar.svelte';
	import DropdownHeader from 'flowbite-svelte/DropdownHeader.svelte';
	import DropdownItem from 'flowbite-svelte/DropdownItem.svelte';
	import ArrowLeftSolid from 'flowbite-svelte-icons/ArrowLeftSolid.svelte';
	import HomeOutline from 'flowbite-svelte-icons/HomeOutline.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import type { User } from '@/lib/types/modelTypes';
	import { routeApi } from '@/lib/util';

	export let withArrowBack: boolean = true;
	export let withHome: boolean = false;
	export let title: string = 'Films';
	export let user: User | null = null;
	export let customBack: string | null = null;
</script>

<Navbar class="fixed top-0 z-50 px-0">
	<div class="flex flex-1 items-center justify-start gap-4">
		{#if withArrowBack}
			<ArrowLeftSolid
				class="h-4 w-4"
				on:click={() => (customBack ? goto(customBack) : history.back())}
			/>
		{/if}
		{#if withHome}
			<HomeOutline
				class="h-4 w-4"
				on:click={async () => {
					await invalidateAll();
					goto('/');
				}}
			/>
		{/if}
	</div>
	<div class="flex-1 font-bold md:text-2xl">
		<h3 class="text-center">{title}</h3>
	</div>
	<div class="flex flex-1 justify-end md:order-2">
		<DarkMode class="me-2" />
		{#if user}
			<div class="flex items-center md:order-2">
				<Avatar id="user-menu" class="object-cover object-center" src={user.avatar ? routeApi(user.avatar) : ''} />
			</div>
			<Dropdown placement="bottom" triggeredBy="#user-menu">
				<DropdownHeader>
					<span class="block truncate text-sm font-medium">{user.nama}</span>
					<span class="block truncate text-sm font-medium">{user.email}</span>
				</DropdownHeader>
				<DropdownItem><a href="/user">Dasboard</a></DropdownItem>
				<DropdownItem><a href="/user/logout">Log out</a></DropdownItem>
			</Dropdown>
		{:else}
			<a
				href="/login"
				class="my-auto rounded-lg border-2 border-primary-700 px-4 py-2 text-xs font-medium text-primary-700 md:text-sm"
				>Sign In</a
			>
		{/if}
	</div>
</Navbar>
