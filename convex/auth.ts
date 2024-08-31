import { convexAuth } from '@convex-dev/auth/server';
import GitHub from '@auth/core/providers/github';
import { Password } from '@convex-dev/auth/providers/password';
import Google from '@auth/core/providers/google';

export const { auth, signIn, signOut, store } = convexAuth({
  providers: [Password, GitHub, Google],
});
