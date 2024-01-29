<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
	import ArrowLeftSolid from 'flowbite-svelte-icons/ArrowLeftSolid.svelte';
</script>

<div class="relative text-gray-700 dark:text-white">
	<Navbar class="fixed top-0 z-50">
		<a href="/film">
			<ArrowLeftSolid class="me-2 h-5 w-5 md:h-8 md:w-8" />
		</a>
		<div class="flex md:order-2">
			<DarkMode class="me-1 md:me-2" />
		</div>
	</Navbar>
	<div class="mt-16">
		<slot />
	</div>
</div>
