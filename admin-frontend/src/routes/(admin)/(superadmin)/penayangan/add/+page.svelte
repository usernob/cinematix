<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Select from 'flowbite-svelte/Select.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import ActionModal from '@/components/actionModal.svelte';
	import { invalidateAll, goto } from '$app/navigation';
	import type { ActionData } from './$types';
	import Modal from 'flowbite-svelte/Modal.svelte';
	import PlusSolid from 'flowbite-svelte-icons/PlusSolid.svelte';
	import SearchOutline from 'flowbite-svelte-icons/SearchOutline.svelte';
	import Search from 'flowbite-svelte/Search.svelte';
	import type { Film } from '@/lib/types/modelTypes';
	import { routeApi } from '@/lib/util';
	import type { ApiResponse } from '@/lib/types/apiResponse';

	export let form: ActionData;

	let selected_film_id: number;

	let formModalState: boolean = false;
	let formModalQuery: string = '';
	let films: Film[];

	let modalState: boolean = false;
	let modalMessage: string = '';
	let modalSuccess: boolean = false;

	const changeSelectedFilmID = (film: Film) => {
		selected_film_id = film.id;
		formModalState = false;
	};

	const findFilm = async (query: string) => {
		const lowercaseQuery: string = query.toLowerCase();
		const req = await fetch(
			routeApi('films/search?' + new URLSearchParams({ q: lowercaseQuery })),
			{
				headers: {
					'Content-Type': 'application/x-www-form-urlencoded',
				},
			}
		);

		const res: ApiResponse<Film[] | null> = await req.json();
		if (res.status === 'ok' && res.data) {
			films = res.data;
		}
	};
</script>

<form
	method="POST"
	enctype="multipart/form-data"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			modalState = true;
			if (result.type === 'success') {
				modalSuccess = true;
				modalMessage = 'Penayangan Berhasil Diupdate';
			} else {
				modalSuccess = false;
				modalMessage = 'Penayangan Gagal Diupdate';
			}
		}}
>
	<div class="flex flex-col items-start gap-4 md:flex-row">
		<div class="w-full flex-1">
			<Label class="mb-4 space-y-2">
				<span>Film ID</span>
				<div class="grid grid-cols-3 items-center gap-2">
					<Input
						class="col-span-2"
						type="text"
						placeholder="Something..."
						size="md"
						value={form?.film_id ?? selected_film_id}
						disabled
						required
					/>
					<Input
						class="col-span-2"
						type="hidden"
						size="md"
						value={form?.film_id ?? selected_film_id}
						name="film_id"
						required
					/>
					<Button size="md" on:click={() => (formModalState = true)}>Cari Film</Button>
				</div>
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Audiotorium ID</span>
				<Select
					name="audiotorium_id"
					size="md"
					placeholder="Pilih Audiotorium"
					class="mt-2"
					items={[
						{ value: 1, name: '1' },
						{ value: 2, name: '2' },
						{ value: 3, name: '3' },
					]}
					value={form?.audiotorium_id ?? ''}
				/>
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Harga</span>
				<Input
					type="text"
					pattern="[0-9]*"
					placeholder="999999"
					size="md"
					value={form?.harga ?? ''}
					name="harga"
					required
				/>
			</Label>
		</div>
		<div class="w-full flex-1">
			<Label class="mb-4 space-y-2">
				<span>Tanggal</span>
				<Input type="date" size="md" value={form?.tanggal ?? ''} name="tanggal" required />
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Mulai</span>
				<Input type="time" size="md" value={form?.mulai ?? ''} name="mulai" required />
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Selesai</span>
				<Input type="time" size="md" value={form?.selesai ?? ''} name="selesai" required />
			</Label>
		</div>
	</div>
	{#if form?.message}
		<div class="text-md my-4 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">{form?.message}</p>
		</div>
	{/if}
	<Button type="submit" class="mt-8">Update</Button>
	<Button
		type="button"
		color="alternative"
		on:click={async () => {
			await invalidateAll();
			await goto('/penayangan');
		}}
		class="mt-8">Batal</Button
	>
	<ActionModal
		bind:modalState
		bind:success={modalSuccess}
		bind:message={modalMessage}
		onConfirm={async () => {
			await invalidateAll();
			if (modalSuccess) {
				await goto('/penayangan');
			}
		}}
	/>
	<Modal
		bind:open={formModalState}
		backdropClass="fixed inset-0 z-[80] bg-gray-900 bg-opacity-50 dark:bg-opacity-80"
		dialogClass="fixed top-0 start-0 end-0 h-modal md:inset-0 md:h-full z-[90] w-full p-4 flex"
		title="Tambah Genre"
		dismissable
	>
		<Label class="mb-4 space-y-2">
			<span>Masukkan Nama Film</span>
			<div class="flex gap-2">
				<Search size="md" bind:value={formModalQuery} />
				<Button class="!p-2.5" on:click={async () => await findFilm(formModalQuery)}>
					<SearchOutline class="h-5 w-5" />
				</Button>
			</div>
		</Label>
		{#if films}
			{#if films.length > 0}
				<div class="flex w-full flex-col divide-y divide-gray-300 dark:divide-gray-500">
					{#each films as film}
						<div class="flex w-full items-center justify-between py-4">
							<p>{film.title}</p>
							<button type="button" on:click={() => changeSelectedFilmID(film)}>
								<PlusSolid class="h-4 w-4 text-primary-700" />
							</button>
						</div>
					{/each}
				</div>
			{:else}
				<p>Film Tidak Ditemukan</p>
			{/if}
		{/if}
	</Modal>
</form>
