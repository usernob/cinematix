<script lang="ts">
	import Navbar from 'flowbite-svelte/Navbar.svelte';
	import NavBrand from 'flowbite-svelte/NavBrand.svelte';
	import NavLi from 'flowbite-svelte/NavLi.svelte';
	import NavUl from 'flowbite-svelte/NavUl.svelte';
	import NavHamburger from 'flowbite-svelte/NavHamburger.svelte';
	import DarkMode from 'flowbite-svelte/DarkMode.svelte';
</script>

<div class="relative text-gray-700 dark:text-white">
	<Navbar class="fixed top-0 z-50">
		<NavBrand href="/">
			<!-- <img src="" class="me-2 h-6 sm:h-9" alt="Logo" /> -->
			<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
				>Flowbite</span
			>
		</NavBrand>
		<div class="flex md:order-2">
			<DarkMode class="me-1 md:me-2" />
			<a
				href="/login"
				class="rounded-lg border-2 border-primary-700 px-4 py-2 text-sm font-medium text-primary-700"
				>Sign In</a
			>
			<NavHamburger class="ms-1 md:ms-3" />
		</div>
		<NavUl>
			<NavLi href="/" active={true}>Home</NavLi>
			<NavLi href="/about">About</NavLi>
		</NavUl>
	</Navbar>
	<slot />
</div>
