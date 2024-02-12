<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import Dropdown from 'flowbite-svelte/Dropdown.svelte';
	import Avatar from 'flowbite-svelte/Avatar.svelte';
	import DropdownHeader from 'flowbite-svelte/DropdownHeader.svelte';
	import DropdownItem from 'flowbite-svelte/DropdownItem.svelte';
	import NavHamburger from 'flowbite-svelte/NavHamburger.svelte';
	import type { User } from '@/lib/types/modelTypes';
	import { routeApi } from '@/lib/util';

	export let title: string = 'Films';
	export let user: User | null = null;
</script>

<Navbar class="fixed top-0 z-50 px-0">
	<NavHamburger
			btnClass="focus:outline-none whitespace-normal rounded-lg focus:ring-2 p-1.5 focus:ring-gray-400 hover:bg-gray-100 dark:hover:bg-gray-600 m-0 mr-3 lg:hidden"
		/>
	<div class="font-bold md:text-2xl flex-1">
		<h3 class="text-center">{title}</h3>
	</div>
	<div class="flex md:order-2 flex-1 justify-end">
		<DarkMode class="me-1 md:me-2" />
		{#if user}
			<div class="flex items-center md:order-2">
        <Avatar id="user-menu" src={user.avatar ? routeApi(user.avatar) : ''}/>
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
				class="rounded-lg border-2 border-primary-700 px-4 py-2 text-sm font-medium text-primary-700"
				>Sign In</a
			>
		{/if}
	</div>
</Navbar>
