import { Tabs, TabList, TabTrigger, TabSlot, type TabTriggerSlotProps } from 'expo-router/ui';
import { Pressable, StyleSheet, Text, View } from 'react-native';
import { AppIcon, type IconName } from '@/components/ui/app-ui';
import { MaxContentWidth, Radius, Spacing } from '@/constants/theme';
import { useAppTheme } from '@/hooks/use-app-theme';

const tabs: { name: string; href: '/' | '/explore' | '/create' | '/nearby' | '/you'; label: string; icon: IconName }[] = [
  { name: 'home', href: '/', label: 'Discover', icon: { ios: 'person.2', android: 'group', web: 'group' } },
  { name: 'explore', href: '/explore', label: 'Search', icon: { ios: 'magnifyingglass', android: 'search', web: 'search' } },
  { name: 'create', href: '/create', label: 'Create', icon: { ios: 'plus.circle', android: 'add_circle', web: 'add_circle' } },
  { name: 'nearby', href: '/nearby', label: 'Nearby', icon: { ios: 'location', android: 'near_me', web: 'near_me' } },
  { name: 'you', href: '/you', label: 'You', icon: { ios: 'person.crop.circle', android: 'account_circle', web: 'account_circle' } },
];
export default function AppTabs() { const { colors } = useAppTheme(); return <Tabs><TabSlot style={{ height: '100%' }} /><TabList asChild><View style={[styles.wrap, { backgroundColor: colors.surface, borderColor: colors.border }]}>{tabs.map(tab => <TabTrigger key={tab.name} name={tab.name} href={tab.href} asChild><TabButton label={tab.label} icon={tab.icon} /></TabTrigger>)}</View></TabList></Tabs>; }
function TabButton({ label, icon, isFocused, ...props }: TabTriggerSlotProps & { label: string; icon: IconName }) { const { colors } = useAppTheme(); return <Pressable {...props} accessibilityLabel={label} style={({ pressed }) => [styles.tab, { backgroundColor: isFocused ? colors.primarySoft : 'transparent', opacity: pressed ? .65 : 1 }]}><AppIcon name={icon} color={isFocused ? colors.primary : colors.textSecondary} /><Text style={[styles.label, { color: isFocused ? colors.primary : colors.textSecondary }]}>{label}</Text></Pressable>; }
const styles = StyleSheet.create({ wrap: { position: 'absolute', bottom: Spacing.md, alignSelf: 'center', width: '95%', maxWidth: MaxContentWidth, flexDirection: 'row', justifyContent: 'space-around', padding: Spacing.sm, borderWidth: StyleSheet.hairlineWidth, borderRadius: Radius.xl, boxShadow: '0 8px 28px rgba(15,23,42,.14)' }, tab: { minWidth: 64, minHeight: 52, borderRadius: Radius.md, alignItems: 'center', justifyContent: 'center', paddingHorizontal: Spacing.sm }, label: { fontSize: 12, marginTop: 2, fontWeight: '700' } });
