import { routeApi } from '@/lib/util';
import type { PageServerLoad } from './$types';
import type { Film, Penayangan } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { fail, type Actions } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, fetch }) => {
	const req = await fetch(routeApi(`penayangan/${params.penayangan_id}`));
	const res: ApiResponse<Penayangan & { film: Film }> = await req.json();

	if (res.status === 'error') {
		throw new Error('Penayangan not found');
	}

	return {
		penayangan: res.data,
		title: 'Edit Penayangan',
		penayangan_id: parseInt(params.penayangan_id),
	};
};

export const actions: Actions = {
	default: async ({ request, fetch, params }) => {
		const body = await request.formData();
		const data = new FormData();
		const film_id = body.get('film_id');
		const audiotorium_id = body.get('audiotorium_id');
		const harga = body.get('harga');
		const tanggal = body.get('tanggal');
		const mulai = body.get('mulai');
		const selesai = body.get('selesai');

		if (!film_id || !audiotorium_id || !harga || !tanggal || !mulai || !selesai) {
			return fail(400, {
				missing: true,
				message: 'All fields are required',
				film_id,
				audiotorium_id,
				harga,
				tanggal,
				mulai,
				selesai,
			});
		}

		data.append('film_id', film_id);
		data.append('audiotorium_id', audiotorium_id);
		data.append('harga', harga);
		data.append('tanggal', tanggal);
		data.append('mulai', mulai);
		data.append('selesai', selesai);

		const req = await fetch(routeApi(`admin/penayangan/${params.penayangan_id}`), {
			method: 'PUT',
			body: data,
		});

		const res: ApiResponse<null> = await req.json();

		if (res.status === 'ok') {
			return {
				success: true,
			};
		} else {
			return fail(500, {
				message: res.message,
				film_id,
				audiotorium_id,
				harga,
				tanggal,
				mulai,
				selesai,
			});
		}
	},
};
