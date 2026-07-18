import { useEffect, useState } from 'react';
import { BackHandler, KeyboardAvoidingView, Modal, Platform, Pressable, ScrollView, StyleSheet, Text, View } from 'react-native';

import { useAuth, type Registration, type VerificationChannel } from '@/auth/AuthContext';
import { AppIcon, BackButton, Button, Card, Field, StatePanel, StatusBanner } from '@/components/ui/app-ui';
import { Radius, Spacing, Type } from '@/constants/theme';
import { useAppTheme } from '@/hooks/use-app-theme';

type Mode = 'login' | 'register' | 'verify';
type Errors = Partial<Record<keyof Registration | 'confirmPassword' | 'code', string>>;

type Country = { code: string; name: string; dial: string; currency: string };
const countries: Country[] = [
  { code: 'NG', name: 'Nigeria', dial: '+234', currency: 'NGN' },
  { code: 'GH', name: 'Ghana', dial: '+233', currency: 'GHS' },
  { code: 'KE', name: 'Kenya', dial: '+254', currency: 'KES' },
  { code: 'ZA', name: 'South Africa', dial: '+27', currency: 'ZAR' },
  { code: 'GB', name: 'United Kingdom', dial: '+44', currency: 'GBP' },
  { code: 'US', name: 'United States', dial: '+1', currency: 'USD' },
  { code: 'CA', name: 'Canada', dial: '+1', currency: 'CAD' },
  { code: 'FR', name: 'France', dial: '+33', currency: 'EUR' },
  { code: 'DE', name: 'Germany', dial: '+49', currency: 'EUR' },
  { code: 'IN', name: 'India', dial: '+91', currency: 'INR' },
];
const localeCountryCode = Intl.DateTimeFormat().resolvedOptions().locale.split('-')[1]?.toUpperCase();
const defaultCountry = countries.find(country => country.code === localeCountryCode) ?? countries[0];
const emptyRegistration: Registration = { full_name: '', username: '', country_code: defaultCountry.code, phone_number: defaultCountry.dial, email: '', date_of_birth: '', password: '' };

export function AuthGate({ children }: React.PropsWithChildren) {
  const { loading, token, login, register, confirmPhone, confirmEmail } = useAuth();
  const { colors } = useAppTheme();
  const [mode, setMode] = useState<Mode>('login');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [registration, setRegistration] = useState(emptyRegistration);
  const [confirmPassword, setConfirmPassword] = useState('');
  const [code, setCode] = useState('');
  const [verificationChannel, setVerificationChannel] = useState<VerificationChannel>('sms');
  const [showPassword, setShowPassword] = useState(false);
  const [countryPickerOpen, setCountryPickerOpen] = useState(false);
  const [error, setError] = useState('');
  const [errors, setErrors] = useState<Errors>({});
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    if (mode === 'login') return;
    const subscription = BackHandler.addEventListener('hardwareBackPress', () => { changeMode(mode === 'verify' ? 'register' : 'login'); return true; });
    return () => subscription.remove();
  }, [mode]);

  if (loading) return <View style={[styles.center, { backgroundColor: colors.background }]}><StatePanel busy icon={{ ios: 'lock.shield', android: 'shield_lock', web: 'shield_lock' }} title="Opening Celebut" message="Restoring your secure session…" /></View>;
  if (token) return children;

  const changeMode = (next: Mode) => { setMode(next); setError(''); setErrors({}); setShowPassword(false); };
  const update = (key: keyof Registration, value: string) => { setRegistration(current => ({ ...current, [key]: value })); setErrors(current => ({ ...current, [key]: undefined })); };
  const selectCountry = (country: Country) => {
    setRegistration(current => {
      const previous = countries.find(item => item.code === current.country_code);
      const localNumber = previous && current.phone_number.startsWith(previous.dial) ? current.phone_number.slice(previous.dial.length) : current.phone_number.replace(/^\+[0-9]{1,3}/, '');
      return { ...current, country_code: country.code, phone_number: `${country.dial}${localNumber}` };
    });
    setErrors(current => ({ ...current, country_code: undefined, phone_number: undefined }));
    setCountryPickerOpen(false);
  };
  const validateRegistration = () => {
    const next: Errors = {};
    if (registration.full_name.trim().length < 3) next.full_name = 'Enter your first and last name.';
    if (!/^[a-zA-Z0-9_]{3,50}$/.test(registration.username.trim())) next.username = 'Use 3–50 letters, numbers, or underscores.';
    if (!/^[A-Za-z]{2,3}$/.test(registration.country_code.trim())) next.country_code = 'Use a 2–3 letter country code, such as NG.';
    if (!/^\+[0-9]{10,15}$/.test(registration.phone_number.trim())) next.phone_number = 'Use international format, such as +2348012345678.';
    if (!/^\S+@\S+\.\S+$/.test(registration.email.trim())) next.email = 'Enter a valid email address.';
    if (!/^\d{4}-\d{2}-\d{2}$/.test(registration.date_of_birth.trim())) next.date_of_birth = 'Use YYYY-MM-DD.';
    if (registration.password.length < 8 || !/[A-Z]/.test(registration.password) || !/[0-9]/.test(registration.password) || !/[^A-Za-z0-9]/.test(registration.password)) next.password = 'Use 8+ characters with an uppercase letter, number, and symbol.';
    if (confirmPassword !== registration.password) next.confirmPassword = 'Passwords do not match.';
    setErrors(next);
    return Object.keys(next).length === 0;
  };
  const submitLogin = async () => { setSubmitting(true); setError(''); try { await login(email, password); } catch (err) { setError(err instanceof Error ? `${err.message}. Check your email and password, then try again.` : 'We could not sign you in. Check your details and try again.'); } finally { setSubmitting(false); } };
  const submitRegistration = async () => { if (!validateRegistration()) return; setSubmitting(true); setError(''); try { const channel = await register({ ...registration, country_code: registration.country_code.trim().toUpperCase(), phone_number: registration.phone_number.trim() }); setVerificationChannel(channel); setEmail(registration.email); setPassword(registration.password); changeMode('verify'); } catch (err) { setError(err instanceof Error ? err.message : 'We could not create your account. Review your details and try again.'); } finally { setSubmitting(false); } };
  const submitCode = async () => { if (!/^\d{6}$/.test(code)) { setErrors({ code: `Enter the 6-digit code sent to your ${verificationChannel === 'email' ? 'email' : 'phone'}.` }); return; } setSubmitting(true); setError(''); try { if (verificationChannel === 'email') await confirmEmail(registration.email, code); else await confirmPhone(registration.phone_number, code); await login(registration.email, registration.password); } catch (err) { setError(err instanceof Error ? `${err.message}. Check the code and try again.` : 'We could not verify this code. Try again.'); } finally { setSubmitting(false); } };

  return <KeyboardAvoidingView behavior={Platform.OS === 'ios' ? 'padding' : undefined} style={[styles.page, { backgroundColor: colors.background }]}>
    <ScrollView keyboardShouldPersistTaps="handled" contentContainerStyle={styles.scroll}>
      <View style={styles.content}>
        <View style={[styles.mark, { backgroundColor: colors.primary }]}><AppIcon name={{ ios: 'sparkles', android: 'celebration', web: 'celebration' }} color={colors.onPrimary} size={28} /></View>
        <Text accessibilityRole="header" style={[styles.brand, { color: colors.text }]}>Celebut</Text>
        <Text style={[styles.promise, { color: colors.textSecondary }]}>Find your people. Celebrate what matters.</Text>
        {mode === 'login' ? <LoginCard colors={colors} email={email} setEmail={setEmail} password={password} setPassword={setPassword} showPassword={showPassword} setShowPassword={setShowPassword} error={error} submitting={submitting} submit={submitLogin} onRegister={() => changeMode('register')} /> : null}
        {mode === 'register' ? <Card style={styles.card}>
          <BackButton onPress={() => changeMode('login')} label="Back to sign in" />
          <Text style={[styles.title, { color: colors.text }]}>Create your account</Text><Text style={[styles.copy, { color: colors.textSecondary }]}>Join a community built around meaningful moments, with privacy and safety by design.</Text>
          {error ? <StatusBanner tone="error" title="Account not created" message={error} /> : null}
          <Field label="Full name *" value={registration.full_name} onChangeText={value => update('full_name', value)} autoComplete="name" textContentType="name" returnKeyType="next" error={errors.full_name} />
          <Field label="Username *" value={registration.username} onChangeText={value => update('username', value)} autoCapitalize="none" autoCorrect={false} returnKeyType="next" helper="Letters, numbers, and underscores only." error={errors.username} />
          <CountrySelector country={countries.find(item => item.code === registration.country_code) ?? defaultCountry} open={() => setCountryPickerOpen(true)} error={errors.country_code} />
          <Field label="Phone number *" value={registration.phone_number} onChangeText={value => update('phone_number', value.replace(/(?!^\+)\D/g, ''))} keyboardType="phone-pad" textContentType="telephoneNumber" autoComplete="tel" placeholder={`${(countries.find(item => item.code === registration.country_code) ?? defaultCountry).dial}8012345678`} helper="The international prefix is selected automatically." error={errors.phone_number} />
          <Field label="Email address *" value={registration.email} onChangeText={value => update('email', value)} autoCapitalize="none" keyboardType="email-address" textContentType="emailAddress" autoComplete="email" returnKeyType="next" error={errors.email} />
          <Field label="Date of birth *" value={registration.date_of_birth} onChangeText={value => update('date_of_birth', value)} keyboardType="numbers-and-punctuation" placeholder="YYYY-MM-DD" helper="You must enter a date in the past." error={errors.date_of_birth} />
          <PasswordField label="Password *" value={registration.password} onChangeText={(value: string) => update('password', value)} show={showPassword} setShow={setShowPassword} error={errors.password} autoComplete="new-password" />
          <Field label="Confirm password *" value={confirmPassword} onChangeText={(value: string) => { setConfirmPassword(value); setErrors(current => ({ ...current, confirmPassword: undefined })); }} secureTextEntry={!showPassword} textContentType="newPassword" autoComplete="new-password" returnKeyType="done" onSubmitEditing={submitRegistration} error={errors.confirmPassword} />
          <Button label="Create account" onPress={submitRegistration} loading={submitting} icon={{ ios: 'person.badge.plus', android: 'person_add', web: 'person_add' }} />
          <View style={styles.switchRow}><Text style={[styles.switchCopy, { color: colors.textSecondary }]}>Already have an account?</Text><Pressable accessibilityRole="button" onPress={() => changeMode('login')} style={({ pressed }) => [styles.link, { opacity: pressed ? .6 : 1 }]}><Text style={[styles.linkText, { color: colors.primary }]}>Sign in</Text></Pressable></View>
        </Card> : null}
        {mode === 'verify' ? <Card style={styles.card}>
          <BackButton onPress={() => changeMode('register')} label="Back to account details" />
          <Text style={[styles.title, { color: colors.text }]}>Verify your {verificationChannel === 'email' ? 'email' : 'phone'}</Text><Text style={[styles.copy, { color: colors.textSecondary }]}>Enter the 6-digit code sent to {verificationChannel === 'email' ? registration.email : registration.phone_number}. This helps keep accounts trustworthy.</Text>
          {error ? <StatusBanner tone="error" title="Code not accepted" message={error} /> : null}
          <Field label="Verification code *" value={code} onChangeText={value => { setCode(value.replace(/\D/g, '').slice(0, 6)); setErrors({}); }} keyboardType="number-pad" textContentType="oneTimeCode" autoComplete={verificationChannel === 'sms' ? 'sms-otp' : 'one-time-code'} maxLength={6} returnKeyType="done" onSubmitEditing={submitCode} error={errors.code} />
          <Button label="Verify and continue" onPress={submitCode} loading={submitting} disabled={code.length !== 6} icon={{ ios: 'checkmark.shield', android: 'verified_user', web: 'verified_user' }} />
          <View style={styles.switchRow}><Pressable accessibilityRole="button" onPress={() => changeMode('register')} style={({ pressed }) => [styles.link, { opacity: pressed ? .6 : 1 }]}><Text style={[styles.linkText, { color: colors.primary }]}>Edit account details</Text></Pressable></View>
        </Card> : null}
        <Text style={[styles.privacy, { color: colors.textSubtle }]}>Celebut never shares your precise location with other people.</Text>
      </View>
    </ScrollView>
    <CountryPicker visible={countryPickerOpen} selected={registration.country_code} close={() => setCountryPickerOpen(false)} select={selectCountry} />
  </KeyboardAvoidingView>;
}

function CountrySelector({ country, open, error }: { country: Country; open: () => void; error?: string }) {
  const { colors } = useAppTheme(); return <View style={styles.selectorField}><Text style={[styles.selectorLabel, { color: colors.text }]}>Country or region *</Text><Pressable accessibilityRole="button" accessibilityLabel={`Country or region, ${country.name}`} accessibilityHint="Opens the country selector" onPress={open} style={({ pressed }) => [styles.selector, { backgroundColor: colors.surfaceRaised, borderColor: error ? colors.danger : colors.border, opacity: pressed ? .72 : 1 }]}><View style={styles.selectorCopy}><Text style={[styles.selectorName, { color: colors.text }]}>{country.name}</Text><Text style={[styles.selectorDetail, { color: colors.textSecondary }]}>{country.code} · {country.currency} · {country.dial}</Text></View><AppIcon name={{ ios: 'chevron.down', android: 'expand-more', web: 'expand-more' }} color={colors.textSecondary} /></Pressable>{error ? <Text accessibilityRole="alert" style={[styles.selectorError, { color: colors.danger }]}>{error}</Text> : null}</View>;
}

function CountryPicker({ visible, selected, close, select }: { visible: boolean; selected: string; close: () => void; select: (country: Country) => void }) {
  const { colors } = useAppTheme(); return <Modal visible={visible} animationType="slide" presentationStyle="pageSheet" onRequestClose={close}><View style={[styles.modal, { backgroundColor: colors.background }]}><View style={styles.modalHeader}><View><Text accessibilityRole="header" style={[styles.modalTitle, { color: colors.text }]}>Choose country</Text><Text style={[styles.modalCopy, { color: colors.textSecondary }]}>This sets your country code, currency, and phone prefix.</Text></View><Pressable accessibilityRole="button" accessibilityLabel="Close country selector" onPress={close} style={({ pressed }) => [styles.close, { backgroundColor: colors.surfaceMuted, opacity: pressed ? .65 : 1 }]}><AppIcon name="close" color={colors.text} /></Pressable></View><ScrollView contentContainerStyle={styles.countryList}>{countries.map(country => <Pressable key={country.code} accessibilityRole="radio" accessibilityState={{ checked: selected === country.code }} onPress={() => select(country)} style={({ pressed }) => [styles.countryOption, { backgroundColor: selected === country.code ? colors.primarySoft : colors.surface, borderColor: colors.border, opacity: pressed ? .7 : 1 }]}><View style={styles.selectorCopy}><Text style={[styles.selectorName, { color: colors.text }]}>{country.name}</Text><Text style={[styles.selectorDetail, { color: colors.textSecondary }]}>{country.code} · {country.currency} · {country.dial}</Text></View>{selected === country.code ? <AppIcon name="check-circle" color={colors.primary} /> : null}</Pressable>)}</ScrollView></View></Modal>;
}

function LoginCard({ colors, email, setEmail, password, setPassword, showPassword, setShowPassword, error, submitting, submit, onRegister }: any) {
  return <Card style={styles.card}><Text style={[styles.title, { color: colors.text }]}>Welcome back</Text><Text style={[styles.copy, { color: colors.textSecondary }]}>Your recommendations and celebrations stay private to your account.</Text>
    {error ? <StatusBanner tone="error" title="Sign-in failed" message={error} /> : null}
    <Field label="Email address" value={email} onChangeText={setEmail} autoCapitalize="none" keyboardType="email-address" textContentType="emailAddress" autoComplete="email" returnKeyType="next" />
    <PasswordField label="Password" value={password} onChangeText={setPassword} show={showPassword} setShow={setShowPassword} autoComplete="current-password" onSubmitEditing={submit} />
    <Button label="Sign in" onPress={submit} loading={submitting} disabled={!email.trim() || !password} icon={{ ios: 'arrow.right', android: 'arrow_forward', web: 'arrow_forward' }} />
    <View style={styles.divider}><View style={[styles.line, { backgroundColor: colors.border }]} /><Text style={[styles.or, { color: colors.textSubtle }]}>New to Celebut?</Text><View style={[styles.line, { backgroundColor: colors.border }]} /></View>
    <Button label="Create an account" variant="secondary" onPress={onRegister} icon={{ ios: 'person.badge.plus', android: 'person_add', web: 'person_add' }} />
  </Card>;
}

function PasswordField({ label, value, onChangeText, show, setShow, error, autoComplete, onSubmitEditing }: any) {
  const { colors } = useAppTheme(); return <View><Field label={label} value={value} onChangeText={onChangeText} secureTextEntry={!show} textContentType={autoComplete === 'new-password' ? 'newPassword' : 'password'} autoComplete={autoComplete} returnKeyType={onSubmitEditing ? 'go' : 'next'} onSubmitEditing={onSubmitEditing} helper={!error && autoComplete === 'new-password' ? '8+ characters with an uppercase letter, number, and symbol.' : undefined} error={error} /><Pressable accessibilityRole="button" accessibilityLabel={show ? 'Hide password' : 'Show password'} hitSlop={10} onPress={() => setShow((value: boolean) => !value)} style={styles.passwordToggle}><AppIcon name={{ ios: show ? 'eye.slash' : 'eye', android: show ? 'visibility_off' : 'visibility', web: show ? 'visibility_off' : 'visibility' }} color={colors.textSecondary} /></Pressable></View>;
}

const styles = StyleSheet.create({
  page: { flex: 1 }, center: { flex: 1 }, scroll: { flexGrow: 1, justifyContent: 'center', padding: Spacing.lg }, content: { width: '100%', maxWidth: 500, alignSelf: 'center', alignItems: 'center', paddingVertical: Spacing.md }, mark: { width: 56, height: 56, borderRadius: Radius.lg, alignItems: 'center', justifyContent: 'center', marginBottom: Spacing.sm }, brand: { fontSize: 34, fontFamily: Platform.select({ ios: 'ui-serif', default: 'serif' }), fontWeight: '700' }, promise: { fontSize: Type.body, textAlign: 'center', marginTop: 4, marginBottom: Spacing.lg }, card: { width: '100%', padding: Spacing.lg }, title: { fontSize: Type.heading, fontWeight: '700' }, copy: { fontSize: Type.body, lineHeight: 24, marginTop: 4, marginBottom: Spacing.md }, passwordToggle: { position: 'absolute', right: 12, top: 34, width: 44, height: 44, alignItems: 'center', justifyContent: 'center' }, privacy: { fontSize: Type.caption, lineHeight: 19, textAlign: 'center', marginTop: Spacing.sm, maxWidth: 360 }, switchRow: { minHeight: 48, marginTop: Spacing.md, flexDirection: 'row', flexWrap: 'wrap', alignItems: 'center', justifyContent: 'center', gap: Spacing.xs }, switchCopy: { fontSize: Type.label }, link: { minHeight: 44, justifyContent: 'center', paddingHorizontal: Spacing.xs }, linkText: { fontSize: Type.label, fontWeight: '800' }, divider: { flexDirection: 'row', alignItems: 'center', gap: Spacing.sm, marginVertical: Spacing.md }, line: { height: StyleSheet.hairlineWidth, flex: 1 }, or: { fontSize: Type.caption }, selectorField: { gap: Spacing.sm, marginBottom: Spacing.md }, selectorLabel: { fontSize: Type.label, fontWeight: '700' }, selector: { minHeight: 56, borderWidth: 1, borderRadius: Radius.md, paddingHorizontal: 14, flexDirection: 'row', alignItems: 'center', gap: Spacing.md }, selectorCopy: { flex: 1 }, selectorName: { fontSize: Type.body, fontWeight: '700' }, selectorDetail: { fontSize: Type.caption, marginTop: 2 }, selectorError: { fontSize: Type.caption, lineHeight: 19 }, modal: { flex: 1, paddingTop: Spacing.lg }, modalHeader: { flexDirection: 'row', alignItems: 'flex-start', paddingHorizontal: Spacing.lg, paddingBottom: Spacing.md, gap: Spacing.md }, modalTitle: { fontSize: Type.heading, fontWeight: '800' }, modalCopy: { fontSize: Type.caption, lineHeight: 19, marginTop: 3 }, close: { width: 48, height: 48, borderRadius: 24, alignItems: 'center', justifyContent: 'center' }, countryList: { padding: Spacing.md, gap: Spacing.sm, paddingBottom: Spacing.xxl }, countryOption: { minHeight: 64, borderWidth: StyleSheet.hairlineWidth, borderRadius: Radius.md, paddingHorizontal: Spacing.md, flexDirection: 'row', alignItems: 'center', gap: Spacing.md },
});
