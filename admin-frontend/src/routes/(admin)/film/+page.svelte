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
	import type { Film } from '@/lib/types/modelTypes';
	import ConfirmationModal from '@/components/confirmationModal.svelte';
	import Modal from 'flowbite-svelte/Modal.svelte';

	export let data: PageData;

	let items: Film[] = [];
	let searchTerm = '';

	let modalDetail: boolean = false;
	let filmDetail: Film | null = null;

	let modalConfirm: boolean = false;
	let deleteId: number;

	const sortKey = writable('id'); // default sort key
	const sortDirection = writable(1); // default sort direction (ascending)
	const sortItems = writable(items.slice()); // make a copy of the items array

	onMount(async () => {
		const req = await fetch(routeApi('admin/all-films'), {
			headers: {
				Authorization: `Bearer ${data.token}`,
			},
		});

		const res: ApiResponse<Film[]> = await req.json();
		if (res.status === 'ok') {
			items = res.data;
			sortItems.set(items.slice());
		}
	});

	const deleteFilm = async (id: number) => {
		const req = await fetch(routeApi(`admin/films/${id}`), {
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
		const sorted = [...$sortItems].sort((a: Film, b: Film): number => {
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
		(item) => item.title.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1
	);

	$: sortItems.set(filteredItems.slice());

	const getDetail = (id: number) => {
		const film = items.find((item) => item.id === id);
		if (!film) {
			return;
		}
		filmDetail = film;
		modalDetail = true;
	};
</script>

<TableSearch
	hoverable={true}
	placeholder="Cari Nama Film"
	bind:inputValue={searchTerm}
	searchClass="relative"
	innerDivClass="py-4 flex justify-between items-center gap-4"
>
	<div slot="header">
		{#if data.user?.role === 'superadmin'}
			<Button size="lg" href="/film/add">Tambah Film</Button>
		{/if}
	</div>
	<TableHead theadClass="cursor-pointer">
		<TableHeadCell on:click={() => sortTable('id')}>ID</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('title')}>Judul</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('rating')}>Rating</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('updated_at')}>Terakhir Diubah</TableHeadCell>
		<TableHeadCell>Aksi</TableHeadCell>
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each $sortItems as item (item.id)}
			<TableBodyRow>
				<TableBodyCell>{item.id}</TableBodyCell>
				<TableBodyCell>{item.title}</TableBodyCell>
				<TableBodyCell>{item.rating}</TableBodyCell>
				<TableBodyCell>{new Date(item.updated_at).toLocaleDateString('id-ID')}</TableBodyCell>
				<TableBodyCell
					class="flex items-center justify-start divide-x divide-gray-700 dark:divide-gray-400"
				>
					<!-- svelte-ignore a11y-invalid-attribute -->
					<a
						role="button"
						href="#"
						class="pr-2 text-primary-500 hover:underline"
						on:click={() => getDetail(item.id)}>Detail</a
					>
					{#if data.user?.role === 'superadmin'}
						<a href="/film/{item.id}" class="px-2 text-primary-500 hover:underline">Edit</a>
						<!-- svelte-ignore a11y-invalid-attribute -->
						<a
							role="button"
							href="#"
							on:click={() => {
								modalConfirm = true;
								deleteId = item.id;
							}}
							class="pl-2 text-primary-500 hover:underline"
						>
							Hapus
						</a>
					{/if}
				</TableBodyCell>
			</TableBodyRow>
		{/each}
	</TableBody>
</TableSearch>

<Modal
	bind:open={modalDetail}
	size="lg"
	dismissable={true}
	title={filmDetail?.title}
	backdropClass="fixed inset-0 z-[80] bg-gray-900 bg-opacity-50 dark:bg-opacity-80"
	dialogClass="fixed top-0 start-0 end-0 h-modal md:inset-0 md:h-full z-[90] w-full p-4 flex"
>
	<div class="flex flex-col gap-4 md:flex-row">
		<img
			src={routeApi(filmDetail?.poster_path ?? '')}
			alt={filmDetail?.title}
			class="aspect-[9.2/13] h-auto w-40 max-w-max rounded-lg object-cover md:w-52"
		/>
		<div class="ml-0 grid flex-1 auto-cols-fr grid-cols-2 gap-4 md:ml-10">
			<p>Judul</p>
			<p>{filmDetail?.title}</p>
			<p>Rating</p>
			<p>{filmDetail?.rating}</p>
			<div class="col-span-2">
				<p>Sinopsis</p>
				<p>{filmDetail?.sinopsis}</p>
			</div>
		</div>
	</div>
</Modal>
<ConfirmationModal
	bind:modalState={modalConfirm}
	message="apakah kamu yakin ingin menghapus film ini?"
	onConfirm={() => deleteFilm(deleteId)}
/>
