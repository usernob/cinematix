<script lang="ts">
	import { PUBLIC_APP_NAME } from '$env/static/public';
	import { routeApi } from '$lib/util';

	const loadData = async () => {
		const res = await fetch(routeApi('films'));
		const data = await res.json();
		return { data: data.data };
	};
</script>

<svelte:head>
	<title>{PUBLIC_APP_NAME.toUpperCase()} - Film</title>
</svelte:head>
<div class="container mt-8 md:mt-16">
	<section>
		{#await loadData()}
			<p>Loading...</p>
		{:then data}
			<h1 class="mb-4 text-lg font-extrabold md:text-3xl">
				{PUBLIC_APP_NAME.toUpperCase()} FEATURED
			</h1>
			<div class="grid auto-cols-fr grid-cols-2 gap-6 md:grid-cols-3 lg:grid-cols-5">
				{#each data.data as img}
					<a href="/film/{img.id}">
						<img
							src={routeApi(img?.poster_path)}
							alt={img?.title}
							class="aspect-[9.2/13] h-auto w-full rounded-lg object-cover shadow-lg md:w-64"
						/>
						<p class="text-md mt-2 font-semibold md:text-xl">{img?.title}</p>
					</a>
				{/each}
			</div>
		{/await}
	</section>
</div>
