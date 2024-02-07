<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import type { PageData } from './$types';

	export let data: PageData;

	let harga = 1000000;
</script>

<div class="min-h-screen">
	<div class="container"></div>
	<div class="fixed bottom-0 w-full bg-slate-50 p-4 dark:bg-gray-800">
		<div class="container flex items-center justify-between">
			<div>
				<h4 class="text-sm font-light md:text-2xl">Total</h4>
				<h3 class="text-lg font-semibold md:text-3xl">
					{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(harga)}
				</h3>
			</div>
			<Button href="/film/{data.filmid}/konfirmasi/{data.penayanganid}" size="lg">Next</Button>
		</div>
	</div>
</div>
