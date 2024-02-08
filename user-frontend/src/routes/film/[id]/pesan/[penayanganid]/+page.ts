import { routeApi } from '@/lib/util';
import { dataFilm, type FilmData } from '@/lib/stores/data';
import type { PageLoad } from './$types';
import type { Kursi, Penayangan, Seat } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';

export const load: PageLoad = async ({ params }) => {
  const res = await fetch(routeApi(`films/${params.id}`));
  const data: ApiResponse<FilmData> = await res.json();
  if (!res.ok) {
    throw new Error(data.message);
  }

  dataFilm.set(data.data)
  const fetchKursi = await fetch(routeApi(`kursi/${params.penayanganid}`));

  type DataKursi = Kursi & {
    seat: Seat[];
  }

  const dataKursi: ApiResponse<DataKursi[]> = await fetchKursi.json();
  if (!res.ok) {
    throw new Error(data.message);
  }

  return {
    dataKursi: dataKursi.data,
    penayangan: data.data?.penayangan as Penayangan[],
    filmid: params.id,
    penayanganid: params.penayanganid,
    title: 'Pilih Seat',
  };
};
