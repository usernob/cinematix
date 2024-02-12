import type { PageServerLoad } from './$types';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { routeApi } from '@/lib/util';
import type { Film, Kursi, Penayangan, Tiket } from '@/lib/types/modelTypes';

export const load: PageServerLoad = async ({ fetch, params }) => {
  const res = await fetch(routeApi(`user/pesanan/${params.tiket_id}`));

  type PenayanganTiket = Penayangan[] &
    {
      tiket: Tiket[] &
      {
        kursi: Kursi[];
      }[];
    }[];

  const data: ApiResponse<Film & { penayangan: PenayanganTiket }> = await res.json();

  return {
    dataFilm: data.data,
    title: 'Detail Tiket',
  };
};

