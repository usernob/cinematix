<script lang="ts">
	import { page } from '$app/stores';
	import Header from '@/components/header.svelte';

	$: title = $page.data?.title || 'Dashboard';

	export let data;

	$: activeUrl = $page.url.pathname;

	const superadminRoute: string[] = [
		'penayangan',
		'film',
		'genre',
		'list-petugas',
		'tambah-petugas',
	];
	const adminRoute: string[] = ['penayangan', 'film', 'genre'];
</script>

<Header
	{activeUrl}
	{title}
	routes={data.user?.role === 'superadmin' ? superadminRoute : adminRoute}
	user={data.user}
/>

<div class="mx-auto mt-16 flex w-full px-4">
	<main class="mx-auto w-full lg:ml-64">
		<slot />
	</main>
</div>
