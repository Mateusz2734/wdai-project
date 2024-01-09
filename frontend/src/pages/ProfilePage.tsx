import { useGetUserOffers } from "../api/offers";
import { OfferCard } from "../components/OfferCard";
import { useAuth } from "../hooks/useAuth";

const Profile = () => {
  const { user } = useAuth();
  const { data } = useGetUserOffers(user?.userId ? user.userId : 0);

  return (
    <div>
      {data?.offers?.map((offer) => (
        <OfferCard key={offer.offerId} offer={offer} editable />
      ))}
    </div>
  );
};

export default Profile;
