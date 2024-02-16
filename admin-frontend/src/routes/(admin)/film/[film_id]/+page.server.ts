import { routeApi } from '@/lib/util';
import type { PageServerLoad } from './$types';
import type { Film, Genre } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import type { Actions } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, fetch }) => {
  const req = await fetch(routeApi(`films/${params.film_id}`));
  const res: ApiResponse<Film & { genres: Genre[] }> = await req.json();

  return {
    film: res.data,
    title: 'Film Edit',
    film_id: parseInt(params.film_id),
  };
};

export const actions: Actions = {
  default: async ({ request, fetch, params }) => {
    const body = await request.formData();
    const req = await fetch(routeApi(`admin/films/${params.film_id}`), {
      method: 'PUT',
      body: body,
    });
    const res: ApiResponse<Film & { genres: Genre[] }> = await req.json();
  }
}
