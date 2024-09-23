import { useParams } from 'next/navigation';

import { Id } from '../../convex/_generated/dataModel';

export const useMemberId = () => {
  const params = useParams();
  // console.log(params);

  return params.memberId as Id<'members'>;
};
