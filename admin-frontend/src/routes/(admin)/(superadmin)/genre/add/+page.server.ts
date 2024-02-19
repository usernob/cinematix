import { routeApi } from '@/lib/util';
import type { PageServerLoad } from './$types';
import type { Genre } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { fail, type Actions } from '@sveltejs/kit';

export const load: PageServerLoad = () => {
	return {
		title: 'Tambah Genre',
	};
};

export const actions: Actions = {
	default: async ({ request, fetch }) => {
		const body = await request.formData();
		const data = new FormData();
		const nama = body.get('nama');

		if (!nama) {
			return fail(400, {
				missing: true,
				message: 'All fields are required',
				nama,
			});
		}

		data.append('nama', nama);

		const req = await fetch(routeApi(`admin/genre`), {
			method: 'POST',
			body: data,
		});

		const res: ApiResponse<Genre> = await req.json();

		if (res.status === 'ok') {
			return {
				success: true,
			};
		} else {
			return fail(500, {
				message: res.message,
			});
		}
	},
};
