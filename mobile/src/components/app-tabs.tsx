import { Icon, Label, NativeTabs, VectorIcon } from 'expo-router/unstable-native-tabs';
import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import { useAppTheme } from '@/hooks/use-app-theme';

export default function AppTabs() {
  const { colors } = useAppTheme();
  return <NativeTabs backgroundColor={colors.surface} indicatorColor={colors.primarySoft} iconColor={{ default: colors.textSecondary, selected: colors.primary }} labelStyle={{ default: { color: colors.textSecondary }, selected: { color: colors.primary, fontWeight: '700' } }}>
    <NativeTabs.Trigger name="index"><Label>Discover</Label><Icon sf={{ default: 'person.2', selected: 'person.2.fill' }} androidSrc={<VectorIcon family={MaterialIcons} name="group" />} /></NativeTabs.Trigger>
    <NativeTabs.Trigger name="explore"><Label>Search</Label><Icon sf="magnifyingglass" androidSrc={<VectorIcon family={MaterialIcons} name="search" />} /></NativeTabs.Trigger>
    <NativeTabs.Trigger name="create"><Label>Create</Label><Icon sf={{ default: 'plus.circle', selected: 'plus.circle.fill' }} androidSrc={<VectorIcon family={MaterialIcons} name="add-circle" />} /></NativeTabs.Trigger>
    <NativeTabs.Trigger name="nearby"><Label>Nearby</Label><Icon sf={{ default: 'location', selected: 'location.fill' }} androidSrc={<VectorIcon family={MaterialIcons} name="near-me" />} /></NativeTabs.Trigger>
    <NativeTabs.Trigger name="you"><Label>You</Label><Icon sf={{ default: 'person.crop.circle', selected: 'person.crop.circle.fill' }} androidSrc={<VectorIcon family={MaterialIcons} name="account-circle" />} /></NativeTabs.Trigger>
  </NativeTabs>;
}
