import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import React from 'react';
import { ActivityIndicator, Platform, Pressable, ScrollView, StyleSheet, Text, TextInput, type TextInputProps, type ViewStyle, View, useWindowDimensions } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';
import { IconSize, MaxContentWidth, Radius, Spacing, Type } from '@/constants/theme';
import { useAppTheme } from '@/hooks/use-app-theme';

export type IconName = string | { ios: string; android: string; web: string };
const materialAliases: Record<string, string> = { shield_lock: 'lock' };
export function AppIcon({ name, color, size = IconSize.md }: { name: IconName; color?: string; size?: number }) {
  const { colors } = useAppTheme();
  const requestedName = typeof name === 'string' ? name : name.android || name.web || name.ios;
  const materialName = materialAliases[requestedName] ?? requestedName.replaceAll('_', '-');
  return <MaterialIcons name={materialName as React.ComponentProps<typeof MaterialIcons>['name']} color={color ?? colors.text} size={size} />;
}

export function Screen({ children, scroll = false, style }: React.PropsWithChildren<{ scroll?: boolean; style?: ViewStyle }>) {
  const { colors } = useAppTheme(); const { width } = useWindowDimensions();
  const inset = width >= 768 ? Spacing.xl : Spacing.md;
  const content = <View style={[styles.content, { paddingHorizontal: inset, maxWidth: MaxContentWidth }, style]}>{children}</View>;
  return <SafeAreaView edges={['top', 'left', 'right']} style={[styles.safe, { backgroundColor: colors.background }]}>{scroll ? <ScrollView keyboardShouldPersistTaps="handled" contentContainerStyle={styles.scroll}>{content}</ScrollView> : content}</SafeAreaView>;
}

export function Heading({ eyebrow, title, description, action }: { eyebrow?: string; title: string; description?: string; action?: React.ReactNode }) {
  const { colors } = useAppTheme();
  return <View style={styles.headingRow}><View style={styles.headingCopy}>{eyebrow ? <Text maxFontSizeMultiplier={1.6} style={[styles.eyebrow, { color: colors.accent }]}>{eyebrow}</Text> : null}<Text accessibilityRole="header" maxFontSizeMultiplier={1.5} style={[styles.title, { color: colors.text }]}>{title}</Text>{description ? <Text maxFontSizeMultiplier={1.8} style={[styles.description, { color: colors.textSecondary }]}>{description}</Text> : null}</View>{action}</View>;
}

export function Card({ children, style, accessible, label }: React.PropsWithChildren<{ style?: ViewStyle; accessible?: boolean; label?: string }>) {
  const { colors } = useAppTheme(); return <View accessible={accessible} accessibilityLabel={label} style={[styles.card, { backgroundColor: colors.surface, borderColor: colors.border }, style]}>{children}</View>;
}

export function Button({ label, onPress, icon, variant = 'primary', loading, disabled, hint }: { label: string; onPress: () => void; icon?: IconName; variant?: 'primary'|'secondary'|'quiet'|'danger'; loading?: boolean; disabled?: boolean; hint?: string }) {
  const { colors } = useAppTheme(); const inactive = disabled || loading;
  const palette = variant === 'primary' ? [colors.primary, colors.onPrimary] : variant === 'danger' ? [colors.dangerSoft, colors.danger] : variant === 'secondary' ? [colors.surfaceMuted, colors.text] : ['transparent', colors.primary];
  return <Pressable accessibilityRole="button" accessibilityLabel={label} accessibilityHint={hint} accessibilityState={{ disabled: !!inactive, busy: !!loading }} disabled={inactive} onPress={onPress} style={({ pressed }) => [styles.button, { backgroundColor: palette[0], borderColor: variant === 'quiet' ? 'transparent' : colors.border, opacity: inactive ? .45 : pressed ? .72 : 1 }]}>{loading ? <ActivityIndicator color={palette[1]} /> : <>{icon ? <AppIcon name={icon} color={palette[1]} size={IconSize.sm} /> : null}<Text maxFontSizeMultiplier={1.5} style={[styles.buttonText, { color: palette[1] }]}>{label}</Text></>}</Pressable>;
}

export function BackButton({ onPress, label = 'Go back' }: { onPress: () => void; label?: string }) {
  const { colors } = useAppTheme();
  return <Pressable accessibilityRole="button" accessibilityLabel={label} accessibilityHint="Returns to the previous screen" hitSlop={4} onPress={onPress} style={({ pressed }) => [styles.backButton, { backgroundColor: colors.surfaceMuted, borderColor: colors.border, opacity: pressed ? .65 : 1 }]}><AppIcon name={{ ios: 'chevron.left', android: 'arrow_back', web: 'arrow_back' }} color={colors.text} /></Pressable>;
}

export function Field({ label, helper, error, ...props }: TextInputProps & { label: string; helper?: string; error?: string }) {
  const { colors } = useAppTheme(); const id = React.useId();
  return <View style={styles.field}><Text nativeID={`${id}-label`} style={[styles.fieldLabel, { color: colors.text }]}>{label}</Text><TextInput accessibilityLabelledBy={`${id}-label`} accessibilityHint={helper} placeholderTextColor={colors.textSubtle} {...props} style={[styles.input, { backgroundColor: colors.surfaceRaised, color: colors.text, borderColor: error ? colors.danger : colors.border }, props.style]} />{error ? <Text accessibilityRole="alert" style={[styles.helper, { color: colors.danger }]}>{error}</Text> : helper ? <Text style={[styles.helper, { color: colors.textSecondary }]}>{helper}</Text> : null}</View>;
}

export function StatePanel({ icon, title, message, action, busy }: { icon: IconName; title: string; message: string; action?: React.ReactNode; busy?: boolean }) {
  const { colors } = useAppTheme(); return <View accessibilityLiveRegion="polite" style={styles.state}><View style={[styles.stateIcon, { backgroundColor: colors.primarySoft }]}>{busy ? <ActivityIndicator color={colors.primary} /> : <AppIcon name={icon} color={colors.primary} size={IconSize.lg} />}</View><Text style={[styles.stateTitle, { color: colors.text }]}>{title}</Text><Text style={[styles.stateMessage, { color: colors.textSecondary }]}>{message}</Text>{action}</View>;
}

export function StatusBanner({ tone, title, message }: { tone: 'success'|'warning'|'error'|'info'; title: string; message: string }) {
  const { colors } = useAppTheme(); const map = tone === 'success' ? [colors.successSoft, colors.success, { ios: 'checkmark.circle', android: 'check_circle', web: 'check_circle' } as IconName] : tone === 'warning' ? [colors.warningSoft, colors.warning, { ios: 'clock', android: 'schedule', web: 'schedule' } as IconName] : tone === 'error' ? [colors.dangerSoft, colors.danger, { ios: 'exclamationmark.triangle', android: 'warning', web: 'warning' } as IconName] : [colors.primarySoft, colors.primary, { ios: 'info.circle', android: 'info', web: 'info' } as IconName];
  return <View accessibilityRole={tone === 'error' ? 'alert' : 'summary'} style={[styles.banner, { backgroundColor: map[0] as string }]}><AppIcon name={map[2] as IconName} color={map[1] as string} /><View style={styles.bannerCopy}><Text style={[styles.bannerTitle, { color: map[1] as string }]}>{title}</Text><Text style={[styles.bannerMessage, { color: colors.text }]}>{message}</Text></View></View>;
}

const styles = StyleSheet.create({
  safe: { flex: 1 }, scroll: { flexGrow: 1, paddingBottom: 112 }, content: { flex: 1, width: '100%', alignSelf: 'center', paddingTop: Spacing.md, paddingBottom: 112 },
  headingRow: { flexDirection: 'row', alignItems: 'flex-start', justifyContent: 'space-between', gap: Spacing.md, marginBottom: Spacing.lg }, headingCopy: { flex: 1 }, eyebrow: { fontSize: Type.caption, fontWeight: '800', letterSpacing: .8, textTransform: 'uppercase', marginBottom: Spacing.xs }, title: { fontSize: Type.title, fontWeight: '700', lineHeight: 35 }, description: { fontSize: Type.body, lineHeight: 24, marginTop: Spacing.sm },
  card: { borderWidth: StyleSheet.hairlineWidth, borderRadius: Radius.lg, padding: Spacing.md, marginBottom: Spacing.md, ...Platform.select({ ios: { shadowColor: '#0B1622', shadowOpacity: .06, shadowRadius: 12, shadowOffset: { width: 0, height: 5 } }, android: { elevation: 2 } }) },
  button: { minHeight: 48, paddingHorizontal: Spacing.md, borderRadius: Radius.md, borderWidth: StyleSheet.hairlineWidth, flexDirection: 'row', gap: Spacing.sm, alignItems: 'center', justifyContent: 'center' }, buttonText: { fontSize: Type.label, fontWeight: '700' },
  backButton: { width: 48, height: 48, borderRadius: 24, borderWidth: StyleSheet.hairlineWidth, alignItems: 'center', justifyContent: 'center', marginBottom: Spacing.md },
  field: { gap: Spacing.sm, marginBottom: Spacing.md }, fieldLabel: { fontSize: Type.label, fontWeight: '700' }, input: { minHeight: 52, borderWidth: 1, borderRadius: Radius.md, paddingHorizontal: 14, fontSize: Type.body }, helper: { fontSize: Type.caption, lineHeight: 19 },
  state: { flex: 1, minHeight: 280, alignItems: 'center', justifyContent: 'center', padding: Spacing.xl, gap: Spacing.sm }, stateIcon: { width: 56, height: 56, borderRadius: 28, alignItems: 'center', justifyContent: 'center', marginBottom: Spacing.sm }, stateTitle: { fontSize: Type.heading, fontWeight: '700', textAlign: 'center' }, stateMessage: { fontSize: Type.body, lineHeight: 24, textAlign: 'center', maxWidth: 430, marginBottom: Spacing.sm },
  banner: { borderRadius: Radius.md, padding: Spacing.md, flexDirection: 'row', alignItems: 'flex-start', gap: Spacing.md, marginBottom: Spacing.md }, bannerCopy: { flex: 1 }, bannerTitle: { fontWeight: '800', fontSize: Type.label }, bannerMessage: { fontSize: Type.caption, lineHeight: 19, marginTop: 2 },
});
