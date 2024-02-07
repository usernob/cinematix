import { routeApi } from '@/lib/util';
import type { PageLoad } from './$types';
import type { Penayangan } from '@/lib/types/modelTypes';

export const load: PageLoad = async ({ params }) => {
  const res = await fetch(routeApi(`/films/${params.id}/`));
  const data = await res.json();
  if (data.status == 'ok') {
    return {
      penayangan: data.data?.penayangan as Penayangan,
      filmid: params.id,
      penayanganid: params.penayanganid,
      title: 'Pilih Seat',
    };
  }
  return {
    penayangan: null,
    filmid: params.id,
    penayanganid: params.penayanganid,
    title: 'Pilih Seat',
  };
};
