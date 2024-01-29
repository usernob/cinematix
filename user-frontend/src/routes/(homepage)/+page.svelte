<script>
	import Carousel from 'flowbite-svelte/Carousel.svelte';
	import AngleRightSolid from 'flowbite-svelte-icons/AngleRightSolid.svelte';
	let images = [
		{
			alt: 'Cosmic timetraveler',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com',
		},
		{
			alt: 'Cristina Gottardi',
			src: 'https://source.unsplash.com/1600x900/?computer',
			title: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com',
		},
		{
			alt: 'Johannes Plenio',
			src: 'https://source.unsplash.com/1600x900/?people',
			title: 'johannes-plenio-RwHv7LgeC7s-unsplash.com',
		},
		{
			alt: 'Jonatan Pie',
			src: 'https://source.unsplash.com/1600x900/?lanscape',
			title: 'jonatan-pie-3l3RwQdHRHg-unsplash.com',
		},
		{
			alt: 'Mark Harpur',
			src: 'https://source.unsplash.com/1600x900/?monk',
			title: 'mark-harpur-K2s_YE031CA-unsplash',
		},
		{
			alt: 'Pietro De Grandi',
			src: 'https://source.unsplash.com/1600x900/?rain',
			title: 'pietro-de-grandi-T7K4aEPoGGk-unsplash',
		},
		{
			alt: 'Sergey Pesterev',
			src: 'https://source.unsplash.com/1600x900/?sea',
			title: 'sergey-pesterev-tMvuB9se2uQ-unsplash',
		},
		{
			alt: 'Solo travel goals',
			src: 'https://source.unsplash.com/1600x900/?city',
			title: 'solotravelgoals-7kLufxYoqWk-unsplash',
		},
	];

	/**
	 * @type {{ title: any; alt: any; }}
	 */
	let image;
</script>

<main class="w-full">
	<Carousel
		{images}
		duration={10000}
		let:Controls
		let:Indicators
		class="min-h-[60vh] w-full rounded-none md:min-h-[90vh]"
		on:change={({ detail }) => (image = detail)}
	>
		<div
			class="absolute inset-x-0 inset-y-16 bottom-0 flex flex-col items-start justify-end bg-gradient-to-t from-gray-800 to-transparent pb-10 dark:from-gray-900"
		>
			<div class="container text-white">
				<h2 class="text-lg font-bold md:text-3xl">{image?.title}</h2>
				<p class="text-sm">By {image?.alt}</p>
			</div>
		</div>
		<Indicators let:Indicator let:selected>
			<Indicator class="h-2 w-2 md:h-3 md:w-3 {selected ? 'opacity-100' : 'opacity-40'}" />
		</Indicators>
		<Controls />
	</Carousel>

	<div class="container mt-8 md:mt-16">
		<section>
			<h1 class="mb-4 text-lg font-extrabold md:text-3xl">SEDANG TAYANG DI BIOSKOP</h1>
			<div
				class="flex min-w-full snap-x snap-mandatory items-start justify-start gap-6 overflow-x-scroll"
			>
				{#each images as img}
					<a class="flex-1 snap-start snap-always" href="/film/{img.alt}">
						<img
							src={img?.src}
							alt={img?.alt}
							class="aspect-[9.2/13] h-auto w-32 max-w-max rounded-md object-cover shadow-lg md:w-64"
						/>
						<p class="text-md mt-2 font-semibold md:text-xl">{img?.alt}</p>
					</a>
				{/each}
				<a class="flex-1 snap-start snap-always" href="/film">
					<div
						class="flex aspect-[9.2/13] h-auto w-32 items-center justify-center gap-2 rounded-lg bg-gray-200/50 object-cover shadow-lg md:w-64"
					>
						<p class="text-sm font-semibold md:text-xl">Lihat Semua</p>
						<AngleRightSolid
							class="h-5 w-5 rounded-full bg-gray-200/60 p-1 md:h-8  md:w-8 md:p-2"
						/>
					</div>
				</a>
			</div>
		</section>
	</div>
</main>
