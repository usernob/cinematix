<script lang="ts">
	import { page } from '$app/stores';
	import { PUBLIC_APP_NAME } from '$env/static/public';
	import Header from '@/components/header.svelte';
	import type { PageData } from './$types';
  
  export let data: PageData;

	$: title = $page.data?.title || 'Films'
</script>

<svelte:head>
	<title>{PUBLIC_APP_NAME.toUpperCase()} - {title}</title>
</svelte:head>
<div class="relative text-gray-700 dark:text-white">
  <Header title={title} withArrowBack={true} withHome={true} user={data.user} />	
	<div class="mt-16">
		<slot />
	</div>
</div>
