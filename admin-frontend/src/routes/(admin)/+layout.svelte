<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { sineIn } from 'svelte/easing';
	import Sidebar from 'flowbite-svelte/Sidebar.svelte';
	import SidebarGroup from 'flowbite-svelte/SidebarGroup.svelte';
	import SidebarItem from 'flowbite-svelte/SidebarItem.svelte';
	import SidebarWrapper from 'flowbite-svelte/SidebarWrapper.svelte';
	import Drawer from 'flowbite-svelte/Drawer.svelte';
	import CloseButton from 'flowbite-svelte/CloseButton.svelte';
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import Dropdown from 'flowbite-svelte/Dropdown.svelte';
	import Avatar from 'flowbite-svelte/Avatar.svelte';
	import DropdownHeader from 'flowbite-svelte/DropdownHeader.svelte';
	import DropdownItem from 'flowbite-svelte/DropdownItem.svelte';
	import NavHamburger from 'flowbite-svelte/NavHamburger.svelte';
	import { routeApi } from '@/lib/util';

	$: title = $page.data?.title || 'Dashboard';

	let transitionParams = {
		x: -320,
		duration: 200,
		easing: sineIn,
	};

	export let data;

	let breakPoint: number = 1024;
	let width: number;
	let backdrop: boolean = false;
	let activateClickOutside = true;
	let drawerHidden: boolean = false;
	$: if (width >= breakPoint) {
		drawerHidden = false;
		activateClickOutside = false;
	} else {
		drawerHidden = true;
		activateClickOutside = true;
	}
	onMount(() => {
		if (width >= breakPoint) {
			drawerHidden = false;
			activateClickOutside = false;
		} else {
			drawerHidden = true;
			activateClickOutside = true;
		}
	});
	const toggleSide = () => {
		if (width < breakPoint) {
			drawerHidden = !drawerHidden;
		}
	};
	const toggleDrawer = () => {
		drawerHidden = false;
		console.log('hiii');
	};
	$: activeUrl = $page.url.pathname;
	$: console.log(activeUrl);

	const superadminRoute: string[] = ['profile'];
	const adminRoute: string[] = ['profile'];
	// let spanClass = 'pl-2 self-center text-md text-gray-900 whitespace-nowrap dark:text-white';
</script>

<svelte:window bind:innerWidth={width} />
<header class="fixed top-0 z-50 mx-auto flex w-full px-4">
	<Navbar fluid class="!px-0 lg:ml-72">
		<NavHamburger onClick={toggleDrawer} class="m-0 p-0 lg:hidden" />
		<div class="font-bold md:text-2xl">
			<h3 class="text-center">{title}</h3>
		</div>
		<div class="flex justify-end md:order-2">
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
					<DropdownItem><a href="/">Dasboard</a></DropdownItem>
					<DropdownItem><a href="/logout">Log out</a></DropdownItem>
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
</header>

<Drawer
	transitionType="fly"
	{backdrop}
	{transitionParams}
	bind:hidden={drawerHidden}
	bind:activateClickOutside
	width="w-64"
	class="overflow-scroll pb-32"
	id="sidebar"
>
	<div class="flex items-center">
		<CloseButton on:click={() => (drawerHidden = true)} class="mb-4 dark:text-white lg:hidden" />
	</div>
	<Sidebar
		asideClass="w-54"
		{activeUrl}
		activeClass="
      flex items-center 
      p-2 text-base font-normal text-primary-900 
      bg-primary-200 dark:bg-primary-700 
      rounded-lg dark:text-white hover:bg-primary-100 dark:hover:bg-gray-700
    "
	>
		<SidebarWrapper divClass="overflow-y-auto py-4 px-3 rounded dark:bg-gray-800">
			<SidebarGroup>
				<SidebarItem label="Dasboard" href="/" on:click={toggleSide} />
				{#if data.user?.role === 'superadmin'}
					{#each superadminRoute as route}
						<SidebarItem
							label={route.replaceAll('_', ' ')}
							class="capitalize"
							href={'/' + route}
							on:click={toggleSide}
							active={activeUrl === `/${route}`}
						/>
					{/each}
				{:else if data.user?.role === 'admin'}
					{#each adminRoute as route}
						<SidebarItem
							label={route.replaceAll('_', ' ')}
							class="capitalize"
							href={'/' + route}
							on:click={toggleSide}
							active={activeUrl === `/${route}`}
						/>
					{/each}
				{/if}
			</SidebarGroup>
		</SidebarWrapper>
	</Sidebar>
</Drawer>

<div class="mx-auto mt-16 flex w-full px-4">
	<main class="mx-auto w-full lg:ml-72">
		<slot />
	</main>
</div>
