<script lang="ts">
	import Table from 'flowbite-svelte/Table.svelte';
	import TableBody from 'flowbite-svelte/TableBody.svelte';
	import TableBodyCell from 'flowbite-svelte/TableBodyCell.svelte';
	import TableBodyRow from 'flowbite-svelte/TableBodyRow.svelte';
	import TableHead from 'flowbite-svelte/TableHead.svelte';
	import TableHeadCell from 'flowbite-svelte/TableHeadCell.svelte';
	import type { PageServerData } from './$types';
	import { routeApi } from '@/lib/util';
	import type { ApiResponse } from '@/lib/types/apiResponse';
	import ActionModal from '@/components/actionModal.svelte';

	export let data: PageServerData;

	let modalState: boolean = false;
	let modalMessage: string = '';
	let modalSuccess: boolean = false;

	const DeleteAdmin = async (id: number) => {
		const req = await fetch(routeApi(`admin/${id}`), {
			method: 'DELETE',
			headers: {
				Authorization: `Bearer ${data.token}`,
			},
		});
		const res: ApiResponse<null> = await req.json();
		modalState = true;
		if (res.status === 'ok') {
			data.users = data.users.filter((item) => item.id !== id);
			modalSuccess = true;
			modalMessage = 'Berhasil Menghapus Admin';
		} else {
			modalSuccess = false;
			modalMessage = 'Gagal Menghapus Admin';
		}
	};
</script>

<Table hoverable={true}>
	<TableHead theadClass="cursor-pointer">
		<TableHeadCell>ID</TableHeadCell>
		<TableHeadCell>Nama</TableHeadCell>
		<TableHeadCell>Email</TableHeadCell>
		{#if data.user?.role === 'superadmin'}
			<TableHeadCell>Aksi</TableHeadCell>
		{/if}
	</TableHead>
	<TableBody tableBodyClass="divide-y">
		{#each data.users as item}
			<TableBodyRow>
				<TableBodyCell>{item.id}</TableBodyCell>
				<TableBodyCell>{item.nama}</TableBodyCell>
				<TableBodyCell>{item.email}</TableBodyCell>
				{#if data.user?.role === 'superadmin'}
					<TableBodyCell
						><a
							href="#"
							class="text-primary-500 hover:underline"
							on:click={() => DeleteAdmin(item.id)}
							data-sveltekit-preload-data="tap">Hapus</a
						></TableBodyCell
					>
				{/if}
			</TableBodyRow>
		{/each}
	</TableBody>
	<ActionModal bind:modalState bind:success={modalSuccess} bind:message={modalMessage} />
</Table>
