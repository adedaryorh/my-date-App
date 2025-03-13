import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class MainBusinessProfile extends StatefulWidget {
  const MainBusinessProfile({
    required this.constraints,
    super.key,
  });

  final BoxConstraints constraints;

  @override
  State<MainBusinessProfile> createState() => _MainBusinessProfileState();
}

class _MainBusinessProfileState extends State<MainBusinessProfile> {
  @override
  Widget build(BuildContext context) {
    return Positioned(
      top: widget.constraints.maxHeight / 6,
      left: 0,
      right: 0,
      child: Container(
        margin: const EdgeInsets.symmetric(horizontal: 30),
        padding: const EdgeInsets.symmetric(horizontal: 30, vertical: 20),
        width: double.maxFinite,
        decoration: const BoxDecoration(
          borderRadius: BorderRadius.all(
            Radius.circular(20),
          ),
          color: Colors.white,
          boxShadow: [
            BoxShadow(
              color: Colors.black12,
              offset: Offset(0, 20),
              blurRadius: 20,
            ),
          ],
        ),
        child: Column(
          children: [
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceBetween,
              children: [
                Image.asset(
                  AppAssets.walletIcon,
                  width: 41,
                  height: 41,
                ),
                InkWell(
                  onTap: () => context.pushNamed(AppRoute.premiumDetail.name),
                  child: Image.asset(
                    AppAssets.requestVerificationIcon,
                    width: 41,
                    height: 41,
                  ),
                ),
              ],
            ),
            Column(
              children: [
                Text(
                  'Ralph Lauren',
                  style: context.textTheme.titleMedium,
                ),
                Text(
                  'We are a global leader in the design, '
                  'marketing and distribution of '
                  'luxury lifestyle products. For more than 50 years, '
                  'our reputation and distinctive image',
                  style: context.textTheme.bodySmall,
                  textAlign: TextAlign.center,
                ),
                const Space(20),
                Container(
                  alignment: Alignment.center,
                  height: 45,
                  width: 247,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(12),
                    color: context.colorScheme.primary,
                  ),
                  child: Text(
                    'Edit Profile',
                    style: context.textTheme.titleMedium
                        ?.copyWith(color: Colors.white),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
