<script lang="ts">
	import { routeApi } from '@/lib/util';
	import { qr } from '@svelte-put/qr/svg';
	import type { PageData } from './$types';

	export let data: PageData;
	$: penayangan = data.dataFilm.penayangan[0];
	$: seats = penayangan.tiket[0].kursi;
</script>

<div class="container">
	<div class="mx-auto flex w-fit flex-col items-center justify-center gap-8">
		<div class="flex flex-col items-center justify-center gap-8 md:flex-row md:items-start">
			<img
				src={routeApi(data.dataFilm.poster_path ?? '')}
				alt="film poster"
				class="aspect-[9.2/13] h-auto w-36 max-w-max rounded-lg object-cover shadow-lg md:w-56"
			/>
			<div>
				<h4 class="mt-4 text-lg font-semibold md:text-2xl">Detail Pesanan</h4>
				<div
					class="text-md grid auto-cols-fr grid-cols-2 gap-x-4 gap-y-2 text-gray-600 dark:text-gray-400"
				>
					<p>Judul</p>
					<p>{data.dataFilm.title}</p>

					<p>Total Harga</p>
					<p>
						{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(
							penayangan.harga * seats.length
						)}
					</p>

					<p>Audiotorium</p>
					<p>{penayangan.audiotorium_id}</p>

					<p>Tanggal</p>
					<p>
						{new Date(penayangan.mulai).toLocaleDateString('id-ID', {
							weekday: 'long',
							year: 'numeric',
							month: 'long',
							day: 'numeric',
						})}
					</p>

					<p>Waktu</p>
					<p>
						{new Date(penayangan.mulai).toLocaleTimeString('en-US', {
							hour: '2-digit',
							minute: '2-digit',
							hour12: false,
						})}
						~
						{new Date(penayangan.selesai).toLocaleTimeString('en-US', {
							hour: '2-digit',
							minute: '2-digit',
							hour12: false,
						})}
					</p>

					<p>Nomor Kursi</p>
					<p>{seats.map((val) => val.nama).join(', ')}</p>
				</div>
			</div>
		</div>
		<svg
			use:qr={{
				data: `${data.user.id}-${penayangan.tiket[0].id}`,
				shape: 'square',

				anchorInnerFill: 'black',
				anchorOuterFill: 'black',
				moduleFill: 'black',
			}}
			class="aspect-square w-64 rounded-lg bg-white p-4 shadow-lg md:w-96"
		/>
		<p class="text-sm">{`${data.user.id}-${penayangan.tiket[0].id}`}</p>
		<p>Berikan Kode QR ini kepada petugas untuk memvalidasi tiket</p>
	</div>
</div>
