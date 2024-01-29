<script lang="ts">
	import Button from 'flowbite-svelte/Button.svelte';
	import Rating from 'flowbite-svelte/Rating.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<main class="min-h-screen w-full">
	<!-- <img src="https://source.unsplash.com/1600x900/?beach" class="h-64 w-full object-cover" /> -->
	<div class="container flex flex-col gap-4 pt-10 md:flex-row md:gap-8">
		<div>
			<img
				src="https://source.unsplash.com/1600x900/?beach"
				class="aspect-[9.2/13] h-auto w-52 max-w-max rounded-lg object-cover shadow-lg md:w-80 lg:w-96"
			/>
		</div>
		<div class="md:max-w-md lg:max-w-lg">
			<h2 class="text-2xl font-bold md:text-4xl">Nama Film</h2>
			<Rating count rating={4.95} id="example-4">
				<span class="mx-1.5 h-1 w-1 rounded-full bg-gray-500 dark:bg-gray-400" />
				<a
					href="/"
					class="text-sm font-medium text-gray-900 underline hover:no-underline dark:text-white"
				>
					73 reviews
				</a>
				<span class="mx-1.5 h-1 w-1 rounded-full bg-gray-500 dark:bg-gray-400" />
				<p>imdb</p>
			</Rating>
			<h3 class="mt-4 text-lg font-semibold md:text-2xl">Sinopsis</h3>
			<p class="font-light md:font-medium">
				After losing her laundry store to a fire incident, a single mom Duk-hee receives a call from
				a bank offering a generous upfront loan so she takes it right away. But soon after, she
				finds out that she has become a victim of a voice phishing scheme. Devastated, Duk-hee turns
				to the police for help, only to be more disappointed that there is no hope of retrieving her
				money. At the moment of despair, Jae-min, the phisher who tricked her, rings her phone once
				again to tip Duk-hee off about the criminal organization that locked him up and coerced him
				to make voice phishing calls from China. With that call, Duk-hee makes up her mind to
				investigate this case for herself and take a plane to China with her friends.
			</p>
		</div>
		<div class="flex-1">
			<h3 class="mt-4 text-lg font-semibold md:text-2xl">Jadwal Tayang</h3>
			<div class="flex flex-col justify-center">
				<div>
					<h3 class="text-md mt-4 font-semibold md:text-xl">Rabu, 18 Maret</h3>
					<div class="flex flex-wrap items-center justify-start gap-2">
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
					</div>
				</div>
				<div>
					<h3 class="text-md mt-4 font-semibold md:text-xl">Rabu, 18 Maret</h3>
					<div class="flex flex-wrap items-center justify-start gap-2">
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
					</div>
				</div>
				<div>
					<h3 class="text-md mt-4 font-semibold md:text-xl">Rabu, 18 Maret</h3>
					<div class="flex flex-wrap items-center justify-start gap-2">
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
						<Button href="/film">18:00</Button>
					</div>
				</div>
			</div>
		</div>
	</div>
</main>
