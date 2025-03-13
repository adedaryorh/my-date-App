import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class MainProfileWidget extends StatefulWidget {
  const MainProfileWidget({
    required this.constraints,
    super.key,
  });

  final BoxConstraints constraints;

  @override
  State<MainProfileWidget> createState() => _MainProfileWidgetState();
}

class _MainProfileWidgetState extends State<MainProfileWidget> {
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
                  'Alice Smith',
                  style: context.textTheme.bodyLarge,
                ),
                Text(
                  'D.O.B: August 21st, 1999',
                  style: context.textTheme.bodyMedium,
                ),
                Text(
                  'Livin’ Crusing, Lavida',
                  style: context.textTheme.bodyMedium
                      ?.copyWith(color: const Color(0xff979797)),
                ),
                const Space(20),
                InkWell(
                  onTap: () => context.pushNamed(AppRoute.pProfileDetails.name),
                  child: Container(
                    alignment: Alignment.center,
                    width: 109,
                    height: 36,
                    decoration: BoxDecoration(
                      color: context.colorScheme.primary,
                      borderRadius: const BorderRadius.all(Radius.circular(50)),
                    ),
                    child: Text(
                      'Edit Profile',
                      style: context.textTheme.bodySmall
                          ?.copyWith(fontWeight: FontWeight.w600),
                    ),
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
