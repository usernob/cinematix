<script lang="ts">
	import { page } from '$app/stores';
</script>

<div class="flex h-screen flex-col items-center justify-center">
	<h1 class="text-5xl font-bold opacity-60">{$page.status} - {$page.error?.message}</h1>
	{#if $page.status === 404}
		<p class="text-xl font-semibold">Wow kamu sudah melewati batas dunia</p>
	{/if}
</div>
