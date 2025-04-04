import { FC } from 'react';

import { fetcher, useApi } from '../hooks/useApi';
import ListNews from '../components/ListNews';
import { INews } from '../interface/INews';

export const NewsAskPage: FC = () => {
  const { data: response, error, isLoading } = useApi<INews[]>(`/ask`, fetcher);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading: {error.message}</div>;

  return <ListNews data={response?.data || []} nextLink={response?.nextLink} />;
};
