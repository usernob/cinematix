import { routeApi } from '@/lib/util';
import { dataFilm, type FilmData } from '@/lib/stores/data';
import type { PageLoad } from './$types';
import type { Penayangan } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';

export const load: PageLoad = async ({ params }) => {
  const res = await fetch(routeApi(`films/${params.id}`));
  const data: ApiResponse<FilmData> = await res.json();
  if (!res.ok) {
    throw new Error(data.message);
  }

  dataFilm.set(data.data)
  return {
    penayangan: data.data?.penayangan as Penayangan[],
    filmid: params.id,
    penayanganid: params.penayanganid,
    title: 'Pilih Seat',
  };
};
