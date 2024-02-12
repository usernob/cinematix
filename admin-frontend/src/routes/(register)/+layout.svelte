<svelte:head>
	<script>
		if ('color-theme' in localStorage) {
			// explicit preference - overrides author's choice
			localStorage.getItem('color-theme') === 'dark'
				? window.document.documentElement.classList.add('dark')
				: window.document.documentElement.classList.remove('dark');
		} else {
			// browser preference - does not overrides
			if (window.matchMedia('(prefers-color-scheme: dark)').matches)
				window.document.documentElement.classList.add('dark');
		}
	</script>
</svelte:head>
<div class="flex h-screen items-center justify-center">
	<div class="m-4 w-full rounded-xl border border-gray-200 p-4 dark:border-gray-600 md:w-1/3">
		<slot />
	</div>
</div>
