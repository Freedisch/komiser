import { useRouter } from 'next/router';
import { useEffect } from 'react';

function Home() {
  const router = useRouter();
  const isBetaFlagOnboardingWizard = true;

  useEffect(() => {
    if (isBetaFlagOnboardingWizard) {
      router.push('/onboarding');
    } else {
      router.push('/dashboard');
    }
  }, []);
}

export default Home;
