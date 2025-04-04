import useSWR, { SWRConfiguration } from 'swr';
import api from '../lib/axios';

export type ApiResponse<T> = {
  data: T;
  submission?: T;
  nextLink?: string;
  message?: string;
  pagination?: {
    total: number;
    page: number;
    perPage: number;
  };
};

export function useApi<T>(
  key: string | null,
  fetcher: (url: string) => Promise<ApiResponse<T>>,
  config?: SWRConfiguration<ApiResponse<T>, Error>
) {
  return useSWR<ApiResponse<T>, Error>(key, fetcher, {
    revalidateOnFocus: false,
    shouldRetryOnError: false,
    ...config
  });
}

// Generic fetcher function
export async function fetcher<T>(url: string): Promise<ApiResponse<T>> {
  const response = await api.get<ApiResponse<T>>(url);
  return response.data;
}
