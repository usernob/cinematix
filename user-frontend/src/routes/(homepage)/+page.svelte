<script>
	import { Carousel } from 'flowbite-svelte';
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
		class="min-h-[60vh] md:min-h-[90vh] w-full rounded-none"
		on:change={({ detail }) => (image = detail)}
	>
		<div
			class="absolute inset-x-0 inset-y-16 bottom-0 flex flex-col items-start justify-end bg-gradient-to-t from-gray-800 to-transparent pb-10 dark:from-gray-900"
		>
			<div class="container text-white">
				<h2 class="text-lg md:text-3xl font-bold">{image?.title}</h2>
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
				<div class="flex items-start justify-start gap-6 overflow-x-scroll min-w-full snap-x snap-mandatory">
					{#each images as img}
						<div class="flex-1 snap-always snap-start">
							<img
								src={img?.src}
								alt={img?.alt}
								class="aspect-[9.2/13] h-auto max-w-max w-32 md:w-64 rounded-md object-cover shadow-lg"
							/>
							<p class="mt-2 text-md font-semibold md:text-xl">{img?.alt}</p>
						</div>
					{/each}
				</div>
		</section>
	</div>
</main>
