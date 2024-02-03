<script lang="ts">
	import Button from 'flowbite-svelte/Button.svelte';
	import Rating from 'flowbite-svelte/Rating.svelte';

	export let data;
	const loadData = async () => {
		const res = await fetch(`http://localhost:8823/films/${data.id}`);
		const resdata = await res.json();
		const firstData = resdata.data[0];

		// let jadwal: Record<string, any> = {}
		// if (firstData) {
		//   firstData.penayangan.forEach((item: { mulai: string | number | Date; }) => {
		//     const dateobj: Date = new Date(item.mulai)
		//     const day = dateobj.getDay()
		//     jadwal[day] = item
		//     console.log(jadwal)
		//   })
		// }
		return { resdata: firstData };
	};
</script>

<main class="min-h-screen w-full">
	{#await loadData()}
		<p>Loading...</p>
	{:then data}
		<!-- <img src="https://source.unsplash.com/1600x900/?beach" class="h-64 w-full object-cover" /> -->
		<div class="container flex flex-col gap-4 pt-10 md:flex-row md:gap-8">
			<div>
				<img
					src={`https://image.tmdb.org/t/p/w500/${data.resdata?.poster_path}`}
					alt={data.resdata?.title}
					class="aspect-[9.2/13] h-auto w-52 max-w-max rounded-lg object-cover shadow-lg md:w-80 lg:w-96"
				/>
			</div>
			<div class="md:max-w-md lg:max-w-lg">
				<h2 class="text-2xl font-bold md:text-4xl">{data.resdata?.title}</h2>
				<Rating count rating={data.resdata?.rating} id="example-4">
					<span class="mx-1.5 h-1 w-1 rounded-full bg-gray-500 dark:bg-gray-400" />
					<p>imdb</p>
				</Rating>
				<h3 class="mt-4 text-lg font-semibold md:text-2xl">Sinopsis</h3>
				<p class="font-light md:font-medium">{data.resdata?.sinopsis}</p>
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
	{/await}
</main>
