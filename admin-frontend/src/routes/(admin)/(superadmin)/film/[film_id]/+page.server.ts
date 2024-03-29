import { routeApi } from '@/lib/util';
import type { PageServerLoad } from './$types';
import type { Film, Genre } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { fail, type Actions } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, fetch }) => {
	const req = await fetch(routeApi(`films/${params.film_id}`));
	const res: ApiResponse<Film & { genre?: Genre[] }> = await req.json();

	return {
		film: res.data,
		title: 'Edit Film',
		film_id: parseInt(params.film_id),
	};
};

export const actions: Actions = {
	default: async ({ request, fetch, params }) => {
		const body = await request.formData();
		const data = new FormData();
		const title = body.get('title');
		const rating = body.get('rating');
		const sinopsis = body.get('sinopsis');
		const tanggal_rilis = body.get('tanggal_rilis');
		const poster = body.get('poster') as File;
		const genres = body.getAll('genres[]');

		if (!title || !sinopsis || !rating || !tanggal_rilis) {
			return fail(400, {
				missing: true,
				message: 'All fields are required',
				title,
				rating,
				sinopsis,
				tanggal_rilis,
			});
		}

		data.append('title', title);
		data.append('rating', rating);
		data.append('sinopsis', sinopsis);
		data.append('tanggal_rilis', tanggal_rilis);

		if (poster.size > 0) {
			data.append('poster', poster);
		}

		if (genres && genres.length > 0) {
			genres.forEach((e) => {
				data.append('genres[]', e as string);
			});
		}

		const req = await fetch(routeApi(`admin/films/${params.film_id}`), {
			method: 'PUT',
			body: data,
		});
		const res: ApiResponse<Film & { genres: Genre[] }> = await req.json();

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
