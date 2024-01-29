<script lang="ts">
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import EyeOutline from 'flowbite-svelte-icons/EyeOutline.svelte';
	import EyeSlashOutline from 'flowbite-svelte-icons/EyeSlashOutline.svelte';

	let show: boolean = false;

	let email: string = '';
	let password: string = '';
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>
<form action="">
	<h2 class="mb-8 text-2xl font-bold">Login</h2>
	<div class="mb-6">
		<Label for="email" class="text-md mb-2 block">Email</Label>
		<Input
			type="email"
			id="email"
			placeholder="john.doe@company.com"
			class="text-md border-2"
			bind:value={email}
			autocomplete="email"
			name="email"
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="Password" class="text-md mb-2 block">Password</Label>
		<Input
			type={show ? 'text' : 'password'}
			id="password"
			placeholder="•••••••••"
			class="text-md border-2"
			bind:value={password}
			name="password"
			required
		>
			<button slot="left" on:click={() => (show = !show)} class="pointer-events-auto">
				{#if !show}
					<EyeOutline class="h-6 w-6" />
				{:else}
					<EyeSlashOutline class="h-6 w-6" />
				{/if}
			</button>
		</Input>
	</div>
	<p class="mb-4">
		Belum punya akun? <a
			href="/signup"
			class="text-primary-700 hover:underline dark:text-primary-600">register sekarang</a
		>
	</p>
	<Button type="submit" disabled={password == '' || email == ''}>Login</Button>
</form>
