<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import NavBrand from 'flowbite-svelte/NavBrand.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import { PUBLIC_APP_NAME } from '$env/static/public';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="relative text-gray-700 dark:text-white">
	<Navbar class="fixed top-0 z-50">
		<NavBrand href="/">
			<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
				>{PUBLIC_APP_NAME.toLocaleUpperCase()}</span
			>
		</NavBrand>
		<div class="flex md:order-2">
			<DarkMode class="me-1 md:me-2" />
			{#if data.user}
				<a href="/user" class="font-semibold px-4 py-2">Hi, {data.user.nama}</a>
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
