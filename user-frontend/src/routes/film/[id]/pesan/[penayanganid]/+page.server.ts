import { routeApi } from '@/lib/util';
import { dataFilm, type FilmData } from '@/lib/stores/data';
import type { PageServerLoad } from './$types';
import type { Kursi, Tiket } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { get } from 'svelte/store';
import { redirect } from '@sveltejs/kit';

export const load: PageServerLoad = async ({ params, cookies }) => {
  const fetchKursi = await fetch(routeApi(`kursi/${params.penayanganid}`));

  console.log(cookies.get("token"))
  if(cookies.get("token") == undefined) {
    throw redirect(302, '/login');
  }

  type DataKursi = Kursi & {
    tiket: Tiket[];
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
      token: cookies.get("token"),
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
    token: cookies.get("token"),
    dataKursi: dataKursi.data,
    filmid: params.id,
    penayanganid: params.penayanganid,
    title: 'Pilih Seat',
  };
};
