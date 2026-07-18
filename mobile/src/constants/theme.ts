import '@/global.css';
import { Platform } from 'react-native';

export const Colors = {
  light: {
    background: '#F7F5F0', surface: '#FFFDF9', surfaceRaised: '#FFFFFF', surfaceMuted: '#EEF2F4',
    text: '#17202A', textSecondary: '#52606D', textSubtle: '#687782', border: '#D8DEE3',
    primary: '#1E3A5F', onPrimary: '#FFFFFF', primarySoft: '#E4EDF5', accent: '#2F7D65', accentSoft: '#E2F2EC',
    danger: '#B42318', dangerSoft: '#FDEAE7', warning: '#8A5A00', warningSoft: '#FFF3D6',
    success: '#176B4D', successSoft: '#E0F3EA', focus: '#2C6AA0', scrim: 'rgba(13,25,38,0.56)',
    backgroundElement: '#EEF2F4', backgroundSelected: '#E4EDF5',
  },
  dark: {
    background: '#0F172A', surface: '#172238', surfaceRaised: '#1D2A42', surfaceMuted: '#223149',
    text: '#F7F9FB', textSecondary: '#C0CAD3', textSubtle: '#A8B5C0', border: '#3A4A60',
    primary: '#A9C7E5', onPrimary: '#102942', primarySoft: '#243E5A', accent: '#70C6A8', accentSoft: '#173F35',
    danger: '#FFB4AB', dangerSoft: '#5A2422', warning: '#FFD58A', warningSoft: '#4D3916',
    success: '#8ED8BA', successSoft: '#194638', focus: '#9CCBFA', scrim: 'rgba(0,0,0,0.68)',
    backgroundElement: '#223149', backgroundSelected: '#243E5A',
  },
} as const;

export type Theme = typeof Colors.light;
export type ThemeColor = keyof Theme;
export const Spacing = { xs: 4, sm: 8, md: 16, lg: 24, xl: 32, xxl: 48, half: 2, one: 4, two: 8, three: 16, four: 24, five: 32, six: 64 } as const;
export const Radius = { sm: 8, md: 12, lg: 18, xl: 26, pill: 999 } as const;
export const Type = { display: 34, title: 28, heading: 22, body: 16, label: 15, caption: 13 } as const;
export const IconSize = { sm: 18, md: 22, lg: 28 } as const;
export const Motion = { quick: 160, standard: 220, deliberate: 280 } as const;
export const Fonts = Platform.select({
  ios: { sans: 'system-ui', serif: 'ui-serif', rounded: 'ui-rounded', mono: 'ui-monospace' },
  default: { sans: 'normal', serif: 'serif', rounded: 'normal', mono: 'monospace' },
  web: { sans: 'var(--font-display)', serif: 'var(--font-serif)', rounded: 'var(--font-rounded)', mono: 'var(--font-mono)' },
});
export const BottomTabInset = Platform.select({ ios: 50, android: 80 }) ?? 0;
export const MaxContentWidth = 760;
