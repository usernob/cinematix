<script lang="ts">
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger, DarkMode } from 'flowbite-svelte';
</script>

<div class="relative text-gray-700 dark:text-white">
	<Navbar class="fixed top-0 z-50">
		<NavBrand href="/">
			<img src="" class="me-2 h-6 sm:h-9" alt="Logo" />
			<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
				>Flowbite</span
			>
		</NavBrand>
		<div class="flex md:order-2">
			<DarkMode class="me-2" />
			<a
				href="/auth/signin"
				class="text-primary-700 rounded-lg border-2 border-primary-700 px-4 py-2 text-sm font-medium"
				>Sign In</a
			>
			<NavHamburger />
		</div>
		<NavUl>
			<NavLi href="/" active={true}>Home</NavLi>
			<NavLi href="/about">About</NavLi>
		</NavUl>
	</Navbar>
	<slot />
</div>
