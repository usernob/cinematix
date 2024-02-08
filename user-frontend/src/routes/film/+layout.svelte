<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import ArrowLeftSolid from 'flowbite-svelte-icons/ArrowLeftSolid.svelte';
	import { page } from '$app/stores';
	import { PUBLIC_APP_NAME } from '$env/static/public';
	let title: string = 'Films';
	$: {
    title = $page.data?.title || 'Films'
  }
</script>

<svelte:head>
	<title>{PUBLIC_APP_NAME.toUpperCase()} - {title}</title>
</svelte:head>
<div class="relative text-gray-700 dark:text-white">
	<Navbar class="fixed top-0 z-50">
		<ArrowLeftSolid class="me-2 h-5 w-5 md:h-8 md:w-8" on:click={() => history.back()} />
		<div class="font-bold md:text-2xl">
			{#key title}
				<p>{title}</p>
			{/key}
		</div>
		<div class="flex md:order-2">
			<DarkMode class="me-1 md:me-2" />
		</div>
	</Navbar>
	<div class="mt-16">
		<slot />
	</div>
</div>
