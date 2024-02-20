<script lang="ts">
	import type { PageServerData } from './$types';
	import CheckCircleOutline from 'flowbite-svelte-icons/CheckCircleOutline.svelte';
	import XCircleOutline from 'flowbite-svelte-icons/XCircleOutline.svelte';
	export let data: PageServerData;
</script>

<div class="container max-w-[40rem]">
	<div class="flex flex-col items-center gap-4">
		{#if data.tiket.length === 0}
			<h4 class="text-3xl font-semibold opacity-40">Tidak ada tiket yang dipesan</h4>
		{/if}
		{#each data.tiket as tiket (tiket.id)}
			<a
				href={tiket.status_pembayaran === 'done'
					? `/user/tiket/${tiket.id}`
					: `/user/konfirmasi-pesanan/${tiket.id}`}
				class="flex w-full items-center justify-between rounded-md border border-gray-200 px-4 py-2 shadow-lg"
			>
				<div>
					<h4 class="text-lg font-semibold">
						{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(
							tiket.total_harga
						)}
					</h4>
					<p class="text-sm opacity-80">
						{new Date(tiket.created_at).toLocaleDateString('id-ID', {
							weekday: 'long',
							year: 'numeric',
							month: 'long',
							day: 'numeric',
							hour: '2-digit',
							minute: '2-digit',
							second: '2-digit',
							hour12: false,
						})}
					</p>
				</div>
				<div>
					{#if tiket.status_pembayaran === 'done'}
						<CheckCircleOutline class="h-6 w-6 text-green-500" />
					{:else}
						<XCircleOutline class="h-6 w-6 text-red-500" />
					{/if}
				</div>
			</a>
		{/each}
	</div>
</div>
