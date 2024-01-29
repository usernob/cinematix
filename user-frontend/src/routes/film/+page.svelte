<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="container mt-8 md:mt-16">
	<section>
		<h1 class="mb-4 text-lg font-extrabold md:text-3xl">SEDANG TAYANG DI BIOSKOP</h1>
		<div class="grid auto-cols-fr grid-cols-2 gap-6 md:grid-cols-3 lg:grid-cols-5">
			{#each data.images as img}
				<a href="/film/{img.alt}">
					<img
						src={img?.src}
						alt={img?.alt}
						class="aspect-[9.2/13] h-auto w-full rounded-lg object-cover shadow-lg md:w-64"
					/>
					<p class="text-md mt-2 font-semibold md:text-xl">{img?.alt}</p>
				</a>
			{/each}
		</div>
	</section>
</div>
