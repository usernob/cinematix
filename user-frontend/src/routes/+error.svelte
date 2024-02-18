<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import { page } from '$app/stores';
</script>

<div class="flex h-screen flex-col items-center justify-center opacity-70">
	<h1 class="text-5xl font-bold">{$page.status} - {$page.error?.message}</h1>
	{#if $page.status === 404}
		<p class="text-xl font-semibold">
			Alamat yang dicari tidak ditemukan di server silahkan kembali ke halaman utama
		</p>
	{/if}
	{#if $page.status === 500}
		<p class="text-xl font-semibold">
			Ada kesalahan di server silahkan coba kembali ke halaman utama
		</p>
	{/if}
	<a
		href="/"
		class="text-xl font-semibold text-primary-500 hover:underline"
		on:click={async () => {
			await invalidateAll();
			await goto('/');
		}}>Halaman Utama -></a
	>
</div>
