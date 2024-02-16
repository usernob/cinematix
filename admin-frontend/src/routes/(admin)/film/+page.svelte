<script lang="ts">
	import Table from 'flowbite-svelte/Table.svelte';
	import TableSearch from 'flowbite-svelte/TableSearch.svelte';
	import TableBody from 'flowbite-svelte/TableBody.svelte';
	import TableBodyCell from 'flowbite-svelte/TableBodyCell.svelte';
	import TableBodyRow from 'flowbite-svelte/TableBodyRow.svelte';
	import TableHead from 'flowbite-svelte/TableHead.svelte';
	import TableHeadCell from 'flowbite-svelte/TableHeadCell.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import { writable } from 'svelte/store';
	import type { PageData } from '../$types';
	import { onMount } from 'svelte';
	import { routeApi } from '@/lib/util';
	import type { ApiResponse } from '@/lib/types/apiResponse';
	import type { Film } from '@/lib/types/modelTypes';

	export let data: PageData;

	let items: Film[] = [];
  let searchTerm = '';

	const sortKey = writable('id'); // default sort key
	const sortDirection = writable(1); // default sort direction (ascending)
	const sortItems = writable(items.slice()); // make a copy of the items array


  onMount(async () => {
    const req = await fetch(routeApi("admin/all-films"), {
      headers: {
        Authorization: `Bearer ${data.token}`,
      }
    })

    const res: ApiResponse<Film[]> = await req.json()
    if (res.status === "ok") {
      items = res.data
      sortItems.set(items.slice())
    }
  })



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
			const aVal = a[key];
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
  $: filteredItems = items.filter((item) => item.title.toLowerCase().indexOf(searchTerm.toLowerCase()) !== -1);
  $: sortItems.set(filteredItems.slice());
</script>

<TableSearch hoverable={true} placeholder="Cari Nama film" bind:inputValue={searchTerm} searchClass="relative" innerDivClass="py-4 flex justify-between items-center gap-4">
  <Button size="lg" slot="header">Tambah Film</Button>
	<TableHead theadClass="cursor-pointer">
		<TableHeadCell on:click={() => sortTable('id')}>ID</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('title')}>Judul</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('rating')}>Rating</TableHeadCell>
		<TableHeadCell on:click={() => sortTable('updated_at')}>Terakhir Diubah</TableHeadCell>
		{#if data.user?.role === 'superadmin'}
			<TableHeadCell>Aksi</TableHeadCell>
		{/if}
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each $sortItems as item}
			<TableBodyRow>
				<TableBodyCell>{item.id}</TableBodyCell>
				<TableBodyCell>{item.title}</TableBodyCell>
				<TableBodyCell>{item.rating}</TableBodyCell>
				<TableBodyCell>{new Date(item.updated_at).toLocaleDateString("id-ID")}</TableBodyCell>
				{#if data.user?.role === 'superadmin'}
					<TableBodyCell><a href="/film/{item.id}" class="text-primary-500 hover:underline">Edit</a></TableBodyCell>
				{/if}
			</TableBodyRow>
		{/each}
	</TableBody>
</TableSearch>
