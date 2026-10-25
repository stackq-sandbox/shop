# Checkout v2

The rewritten checkout lives behind the `new_checkout` flag. It is off by default
and is turned on for production by `config/env/prod/app.yaml`.

## Request flow

1. `POST /checkout/v2` decodes an order.
2. The shipping address is validated for the destination country.
3. The order is priced: volume discount, coupon, state sales tax, shipping.
4. The payment provider authorizes the total and the charge is captured.

`POST /checkout/preview` runs steps 1 to 3 and returns the quote without charging.

## Tax rates

Rates are flat state-level fixtures.

| State | Rate (bps) |
| ----- | ---------- |
| Alabama (AL) | 400 |
| Alaska (AK) | 0 |
| Arizona (AZ) | 560 |
| Arkansas (AR) | 650 |
| California (CA) | 725 |
| Colorado (CO) | 290 |
| Connecticut (CT) | 635 |
| Delaware (DE) | 0 |
| Florida (FL) | 600 |
| Georgia (GA) | 400 |
| Hawaii (HI) | 400 |
| Idaho (ID) | 600 |
| Illinois (IL) | 625 |
| Indiana (IN) | 700 |
| Iowa (IA) | 600 |
| Kansas (KS) | 650 |
| Kentucky (KY) | 600 |
| Louisiana (LA) | 445 |
| Maine (ME) | 550 |
| Maryland (MD) | 600 |
| Massachusetts (MA) | 625 |
| Michigan (MI) | 600 |
| Minnesota (MN) | 688 |
| Mississippi (MS) | 700 |
| Missouri (MO) | 423 |
| Montana (MT) | 0 |
| Nebraska (NE) | 550 |
| Nevada (NV) | 685 |
| New Hampshire (NH) | 0 |
| New Jersey (NJ) | 663 |
| New Mexico (NM) | 513 |
| New York (NY) | 400 |
| North Carolina (NC) | 475 |
| North Dakota (ND) | 500 |
| Ohio (OH) | 575 |
| Oklahoma (OK) | 450 |
| Oregon (OR) | 0 |
| Pennsylvania (PA) | 600 |
| Rhode Island (RI) | 700 |
| South Carolina (SC) | 600 |
| South Dakota (SD) | 450 |
| Tennessee (TN) | 700 |
| Texas (TX) | 625 |
| Utah (UT) | 485 |
| Vermont (VT) | 600 |
| Virginia (VA) | 430 |
| Washington (WA) | 650 |
| West Virginia (WV) | 600 |
| Wisconsin (WI) | 500 |
| Wyoming (WY) | 400 |

## Carriers

| Key | Name | Base (cents) | Per kg (cents) | Transit days |
| --- | ---- | ------------ | -------------- | ------------ |
| ground | Ground | 499 | 120 | 5 |
| express | Express | 1299 | 240 | 2 |
| overnight | Overnight | 2499 | 480 | 1 |
| economy | Economy | 349 | 90 | 8 |
| freight | Freight | 4999 | 60 | 10 |
| pickup | Store pickup | 0 | 0 | 3 |

Each zone beyond the first adds 15% to the carrier price.

## Payment providers

| Name | Limit (cents) | Fee (bps) | Fixed fee (cents) |
| ---- | ------------- | --------- | ----------------- |
| card | 5000000 | 290 | 30 |
| paypal | 2500000 | 349 | 49 |
| applepay | 1000000 | 250 | 20 |
| klarna | 1000000 | 329 | 0 |
| sepa | 10000000 | 80 | 35 |
| ach | 10000000 | 80 | 0 |

## Rollback

Turn `new_checkout` off. The v2 tables are additive and can stay in place.
