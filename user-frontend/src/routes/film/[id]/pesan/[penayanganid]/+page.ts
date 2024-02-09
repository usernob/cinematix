import { routeApi } from '@/lib/util';
import { dataFilm, type FilmData } from '@/lib/stores/data';
import type { PageLoad } from './$types';
import type { Kursi, Penayangan, Seat } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { get } from 'svelte/store';

export const load: PageLoad = async ({ params }) => {
  const fetchKursi = await fetch(routeApi(`kursi/${params.penayanganid}`));

  type DataKursi = Kursi & {
    seat: Seat[];
  };

  const dataKursi: ApiResponse<DataKursi[]> = await fetchKursi.json();
  if (!fetchKursi.ok) {
    throw new Error(dataKursi.message);
  }

  if (get(dataFilm)) {
    // update list penayangan with the only selected penayangan
    dataFilm.update((val) => {
      if (!val) return val;

      const selected = val.penayangan.find(
        (penayangan) => penayangan.id === Number(params.penayanganid)
      );

      if (selected) {
        val.penayangan = [selected];
      }

      return val;
    });

    return {
      dataKursi: dataKursi.data,
      filmid: params.id,
      penayanganid: params.penayanganid,
      title: 'Pilih Seat',
    };
  }

  const res = await fetch(routeApi(`films/${params.id}/${params.penayanganid}`));
  const data: ApiResponse<FilmData> = await res.json();
  if (!res.ok) {
    throw new Error(data.message);
  }

  dataFilm.set(data.data);

  return {
    dataKursi: dataKursi.data,
    filmid: params.id,
    penayanganid: params.penayanganid,
    title: 'Pilih Seat',
  };
};
