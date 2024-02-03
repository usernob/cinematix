<script lang="ts">
	const loadData = async () => {
		const res = await fetch('http://localhost:8823/films');
		const data = await res.json();
		return { data: data.data };
	};
</script>

<div class="container mt-8 md:mt-16">
	<section>
		{#await loadData()}
			<p>Loading...</p>
		{:then data}
			<h1 class="mb-4 text-lg font-extrabold md:text-3xl">SEDANG TAYANG DI BIOSKOP</h1>
			<div class="grid auto-cols-fr grid-cols-2 gap-6 md:grid-cols-3 lg:grid-cols-5">
				{#each data.data as img}
					<a href="/film/{img.id}">
						<img
							src={`https://image.tmdb.org/t/p/w500/${img?.poster_path}`}
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
