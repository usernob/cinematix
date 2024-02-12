<script lang="ts">
	import { page } from '$app/stores';
	import Header from '@/components/header.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	$: title = $page.data?.title || 'Dashboard';
	$: arrowback = $page.data?.arrowback || true;
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>
<div>
	<Header {title} withHome={true} withArrowBack={arrowback} customBack="/user" user={data.user} />
	<div class="mt-16">
		<slot />
	</div>
</div>
