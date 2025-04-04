import { FC } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import { INews } from '../interface/INews';
import ListNews from '../components/ListNews';

export const NewsPage: FC = () => {
  const { data: response, error, isLoading } = useApi<INews[]>('/news', fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return <ListNews data={response?.data || []} nextLink={response?.nextLink} />;
};
