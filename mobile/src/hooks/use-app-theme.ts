import { useColorScheme } from 'react-native';
import { Colors } from '@/constants/theme';

export function useAppTheme() {
  const scheme = useColorScheme();
  const mode = scheme === 'dark' ? 'dark' : 'light';
  return { colors: Colors[mode], mode };
}
