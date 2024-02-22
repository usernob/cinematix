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
	import type { Genre } from '@/lib/types/modelTypes';
	import ConfirmationModal from '@/components/confirmationModal.svelte';

	export let data: PageData;

	let items: Genre[] = [];
	let searchTerm = '';

	let modalConfirm: boolean = false;
	let deleteId: number;
	const sortKey = writable('id'); // default sort key
	const sortDirection = writable(1); // default sort direction (ascending)
	const sortItems = writable(items.slice()); // make a copy of the items array

	onMount(async () => {
		const req = await fetch(routeApi('genre'), {
			headers: {
				Authorization: `Bearer ${data.token}`,
			},
		});

		const res: ApiResponse<Genre[]> = await req.json();
		if (res.status === 'ok') {
			items = res.data;
			sortItems.set(items.slice());
		}
	});

	const deleteGenre = async (id: number) => {
		const req = await fetch(routeApi(`admin/genre/${id}`), {
			method: 'DELETE',
			headers: {
				Authorization: `Bearer ${data.token}`,
			},
		});
		const res = await req.json();
		if (res.status === 'ok') {
			items = items.filter((item) => item.id !== id);
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
		const sorted = [...$sortItems].sort((a: Genre, b: Genre): number => {
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
		(item) => item.nama.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1
	);

	$: sortItems.set(filteredItems.slice());
</script>

<TableSearch
	hoverable={true}
	placeholder="Cari Nama Genre"
	bind:inputValue={searchTerm}
	searchClass="relative"
	innerDivClass="py-4 flex justify-between items-center gap-4"
>
	<div slot="header">
		{#if data.user?.role === 'superadmin'}
			<Button size="lg" href="/genre/add">Tambah Genre</Button>
		{/if}
	</div>
	<TableHead theadClass="cursor-pointer">
		<TableHeadCell on:click={() => sortTable('id')}>ID</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('nama')}>Nama</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('updated_at')}>Terakhir Diubah</TableHeadCell>
		{#if data.user?.role === 'superadmin'}
			<TableHeadCell>Aksi</TableHeadCell>
		{/if}
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each $sortItems as item (item.id)}
			<TableBodyRow>
				<TableBodyCell>{item.id}</TableBodyCell>
				<TableBodyCell>{item.nama}</TableBodyCell>
				<TableBodyCell>{new Date(item.updated_at).toLocaleDateString('id-ID')}</TableBodyCell>
				{#if data.user?.role === 'superadmin'}
					<TableBodyCell
						class="flex items-center justify-start divide-x divide-gray-700 dark:divide-gray-400"
					>
						<a href="/genre/{item.id}" class="pr-2 text-primary-500 hover:underline">Edit</a>
						<!-- svelte-ignore a11y-invalid-attribute -->
						<a
							href="#"
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
	message="apakah kamu yakin ingin menghapus genre ini?"
	onConfirm={() => deleteGenre(deleteId)}
/>
