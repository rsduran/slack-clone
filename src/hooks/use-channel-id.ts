import { useParams } from 'next/navigation';

import { Id } from '../../convex/_generated/dataModel';

export const useChannelId = () => {
  const params = useParams();
  // console.log(params);

  return params.channelId as Id<'channels'>;
};
