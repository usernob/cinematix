<script lang="ts">
	import { page } from '$app/stores';
	import Header from '@/components/header.svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	$: title = $page.data?.title || 'User';
	$: arrowback = $page.data?.arrowback || true;
</script>

<div>
	<Header {title} withHome={true} withArrowBack={arrowback} user={data.user} />
	<div class="mt-16">
		<slot />
	</div>
</div>
