<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import NavBrand from 'flowbite-svelte/NavBrand.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import Dropdown from 'flowbite-svelte/Dropdown.svelte';
	import Avatar from 'flowbite-svelte/Avatar.svelte';
	import DropdownHeader from 'flowbite-svelte/DropdownHeader.svelte';
	import DropdownItem from 'flowbite-svelte/DropdownItem.svelte';
	import { PUBLIC_APP_NAME } from '$env/static/public';
	import type { PageData } from './$types';
	import { routeApi } from '@/lib/util';

	export let data: PageData;
</script>

<div class="relative text-gray-700 dark:text-white">
	<Navbar class="fixed top-0 z-50 px-0">
		<NavBrand href="/">
			<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
				>{PUBLIC_APP_NAME.toLocaleUpperCase()}</span
			>
		</NavBrand>
		<div class="flex md:order-2">
			<DarkMode class="me-2" />
			{#if data.user}
				<div class="flex items-center md:order-2">
					<Avatar id="user-menu" src={data.user.avatar ? routeApi(data.user.avatar) : ''} />
				</div>
				<Dropdown placement="bottom" triggeredBy="#user-menu">
					<DropdownHeader>
						<span class="block truncate text-sm font-medium">{data.user.nama}</span>
						<span class="block truncate text-sm font-medium">{data.user.email}</span>
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
	<slot />
</div>
