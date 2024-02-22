<script lang="ts">
	import TableSearch from 'flowbite-svelte/TableSearch.svelte';
	import TableBody from 'flowbite-svelte/TableBody.svelte';
	import TableBodyCell from 'flowbite-svelte/TableBodyCell.svelte';
	import TableBodyRow from 'flowbite-svelte/TableBodyRow.svelte';
	import TableHead from 'flowbite-svelte/TableHead.svelte';
	import TableHeadCell from 'flowbite-svelte/TableHeadCell.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import { writable } from 'svelte/store';
	import type { PageData } from './$types';
	import { onMount } from 'svelte';
	import { routeApi } from '@/lib/util';
	import type { ApiResponse } from '@/lib/types/apiResponse';
	import type { Penayangan, Film } from '@/lib/types/modelTypes';
	import ConfirmationModal from '@/components/confirmationModal.svelte';
	import ActionModal from '@/components/actionModal.svelte';

	export let data: PageData;

	type FilmPenayangan = {
		id: number;
		film_id: number;
		film_title: string;
		tanggal: string;
		mulai: string;
		selesai: string;
		harga: number;
		audiotorium_id: number;
		updated_at: string;
	};

	let items: FilmPenayangan[] = [];
	let searchTerm = '';

	let modalConfirm: boolean = false;
	let deleteId: number;

	let modalAction: boolean = false;
	let modalSuccess: boolean = false;
	let modalActionMessage: string = '';

	const sortKey = writable('id'); // default sort key
	const sortDirection = writable(1); // default sort direction (ascending)
	const sortItems = writable(items.slice()); // make a copy of the items array

	onMount(async () => {
		const req = await fetch(routeApi('penayangan'));

		const res: ApiResponse<Penayangan[] & { film: Film }[]> = await req.json();
		if (res.status === 'ok') {
			for (const penayangan of res.data) {
				const filmPenayangan: FilmPenayangan = {
					id: penayangan.id,
					film_id: penayangan.film.id,
					film_title: penayangan.film.title,
					tanggal: penayangan.tanggal,
					mulai: penayangan.mulai,
					selesai: penayangan.selesai,
					harga: penayangan.harga,
					audiotorium_id: penayangan.audiotorium_id,
					updated_at: penayangan.updated_at,
				};
				items = [...items, filmPenayangan];
			}
			sortItems.set(items.slice());
		}
	});

	const deletePenayangan = async (id: number) => {
		const req = await fetch(routeApi(`admin/penayangan/${id}`), {
			method: 'DELETE',
			headers: {
				Authorization: `Bearer ${data.token}`,
			},
		});
		const res = await req.json();
		modalAction = true;
		if (res.status === 'ok') {
			items = items.filter((item) => item.id !== id);
			filteredItems = filteredItems.filter((item) => item.id === id);
			modalSuccess = true;
			modalActionMessage = 'Penayangan berhasil dihapus';
		} else {
			modalSuccess = false;
			modalActionMessage = res.message;
		}
	};

	// Define a function to sort the items
	const sortTable = (key: string) => {
		// If the same key is clicked, reverse the sort direction
		if ($sortKey === key) {
			sortDirection.update((val) => -val);
		} else {
			sortKey.set(key);
			sortDirection.set(1);
		}
	};

	$: {
		const key = $sortKey;
		const direction = $sortDirection;
		const sorted = [...$sortItems].sort((a: FilmPenayangan, b: FilmPenayangan): number => {
			// @ts-ignore
			const aVal = a[key];
			// @ts-ignore
			const bVal = b[key];
			if (aVal < bVal) {
				return -direction;
			} else if (aVal > bVal) {
				return direction;
			}
			return 0;
		});
		sortItems.set(sorted);
	}

	$: filteredItems = items.filter(
		(item) => item.film_title.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1
	);

	$: sortItems.set(filteredItems.slice());
</script>

<TableSearch
	hoverable={true}
	placeholder="Cari Berdasarkan Title Film"
	bind:inputValue={searchTerm}
	searchClass="relative"
	innerDivClass="py-4 flex justify-between items-center gap-4"
>
	<div slot="header">
		{#if data.user?.role === 'superadmin'}
			<Button size="lg" href="/penayangan/add">Tambah Penayangan</Button>
		{/if}
	</div>
	<TableHead theadClass="cursor-pointer">
		<TableHeadCell on:click={() => sortTable('id')}>ID</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('tanggal')}>Tanggal</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('mulai')}>Mulai</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('selesai')}>Selesai</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('harga')}>Harga</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('audiotorium_id')}>Audiotorium ID</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('film_title')}>Film Title</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('updated_at')}>Terakhir Diubah</TableHeadCell>
		{#if data.user?.role === 'superadmin'}
			<TableHeadCell>Aksi</TableHeadCell>
		{/if}
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each $sortItems as item (item.id)}
			<TableBodyRow>
				<TableBodyCell>{item.id}</TableBodyCell>
				<TableBodyCell
					>{new Date(item.tanggal).toLocaleDateString('id-ID', {
						weekday: 'short',
						month: 'short',
						day: 'numeric',
						year: 'numeric',
					})}</TableBodyCell
				>
				<TableBodyCell
					>{new Date(item.mulai).toLocaleTimeString('en-US', {
						hour: '2-digit',
						minute: '2-digit',
						hour12: false,
					})}</TableBodyCell
				>
				<TableBodyCell
					>{new Date(item.selesai).toLocaleTimeString('en-US', {
						hour: '2-digit',
						minute: '2-digit',
						hour12: false,
					})}</TableBodyCell
				>
				<TableBodyCell
					>{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(
						item.harga
					)}</TableBodyCell
				>
				<TableBodyCell>{item.audiotorium_id}</TableBodyCell>
				<TableBodyCell>{item.film_title}</TableBodyCell>
				<TableBodyCell>{new Date(item.updated_at).toLocaleDateString('id-ID')}</TableBodyCell>
				{#if data.user?.role === 'superadmin'}
					<TableBodyCell
						class="flex items-center justify-start divide-x divide-gray-700 dark:divide-gray-400"
					>
						<a href="/penayangan/{item.id}" class="pr-2 text-primary-500 hover:underline">Edit</a>
						<!-- svelte-ignore a11y-invalid-attribute -->
						<a
							role="button"
							href="javascript:void(0)"
							on:click={() => {
								modalConfirm = true;
								deleteId = item.id;
							}}
							class="pl-2 text-primary-500 hover:underline"
						>
							Hapus
						</a>
					</TableBodyCell>
				{/if}
			</TableBodyRow>
		{/each}
	</TableBody>
</TableSearch>

<ConfirmationModal
	bind:modalState={modalConfirm}
	message="apakah kamu yakin ingin menghapus penayangan ini?"
	onConfirm={() => deletePenayangan(deleteId)}
/>

<ActionModal
	bind:modalState={modalAction}
	bind:message={modalActionMessage}
	bind:success={modalSuccess}
/>
