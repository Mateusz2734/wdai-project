import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosError } from "axios";

import {
  ApiError,
  GetOffersResponse,
  PostOfferResponse,
  DeleteOfferResponse,
  GetOfferResponse,
  PostOfferPayload,
} from "../types/types";
import { useAuth } from "../hooks/useAuth";
import api from "./api";

export function useGetOffers() {
  const { token } = useAuth();

  const config = { headers: { Authorization: `Bearer ${token}` } };

  return useQuery<GetOffersResponse, ApiError>({
    queryKey: ["offers"],
    queryFn: async () => {
      try {
        const { data } = await api.get<GetOffersResponse>("/offers", config);
        return data;
      } catch (error) {
        const e = error as AxiosError<ApiError>;
        return Promise.reject(e.response?.data);
      }
    },
    throwOnError: false,
  });
}

export function useGetUserOffers(userId: number) {
  const { token } = useAuth();

  const config = { headers: { Authorization: `Bearer ${token}` } };

  return useQuery<GetOffersResponse, ApiError>({
    queryKey: ["offers"],
    queryFn: async () => {
      try {
        const { data } = await api.get<GetOffersResponse>(
          `/users/${userId}/offers`,
          config
        );
        return data;
      } catch (error) {
        const e = error as AxiosError<ApiError>;
        return Promise.reject(e.response?.data);
      }
    },
    throwOnError: false,
  });
}

export function useGetOffer(offerId: number) {
  const { token } = useAuth();

  const config = { headers: { Authorization: `Bearer ${token}` } };

  return useQuery<GetOfferResponse, ApiError>({
    queryKey: ["offers"],
    queryFn: async () => {
      try {
        const { data } = await api.get<GetOfferResponse>(
          `/offers/${offerId}`,
          config
        );
        return data;
      } catch (error) {
        const e = error as AxiosError<ApiError>;
        return Promise.reject(e.response?.data);
      }
    },
    throwOnError: false,
  });
}

export function useDeleteOffer(offerId: number) {
  const { token } = useAuth();

  const config = { headers: { Authorization: `Bearer ${token}` } };

  return useMutation<DeleteOfferResponse, ApiError>({
    mutationKey: ["offers"],
    mutationFn: async () => {
      try {
        const { data } = await api.delete<DeleteOfferResponse>(
          `/offers/${offerId}`,
          config
        );
        return data;
      } catch (error) {
        const e = error as AxiosError<ApiError>;
        return Promise.reject(e.response?.data);
      }
    },
    throwOnError: false,
  });
}

export function usePostOffer(payload: PostOfferPayload) {
  const { token } = useAuth();

  const config = { headers: { Authorization: `Bearer ${token}` } };

  return useMutation<PostOfferResponse, ApiError>({
    mutationKey: ["offers"],
    mutationFn: async () => {
      try {
        const { data } = await api.post<PostOfferResponse>(
          "/offers",
          payload,
          config
        );
        return data;
      } catch (error) {
        const e = error as AxiosError<ApiError>;
        return Promise.reject(e.response?.data);
      }
    },
    throwOnError: false,
  });
}
